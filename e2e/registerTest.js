import http from "k6/http";
import { check, group } from "k6";
import { randomString } from "https://jslib.k6.io/k6-utils/1.2.0/index.js";

function generateRandomEmail() {
  const randomPrefix = randomString(8);
  return `${randomPrefix}@example.com`;
}

function generateTestUser() {
  return {
    age: 17,
    bio: randomString(20),
    email: generateRandomEmail(),
    location: randomString(8),
    name: randomString(8),
    password: randomString(10),
    profile_pic_url: "https://example.com/profile.jpg",
  };
}

export function TestRegistration() {
  const baseUrl = "http://localhost:8080";
  let validUser;

  group("User Registration Scenarios", function () {
    group("Positive Registration", () => {
      validUser = generateTestUser();
      const registrationPayload = JSON.stringify(validUser);
      const registrationParams = {
        headers: {
          "Content-Type": "application/json",
        },
      };

      const registrationResponse = http.post(
        `${baseUrl}/register`,
        registrationPayload,
        registrationParams
      );
      check(registrationResponse, {
        "registration status is success": (r) => r.json().status === "success",
        "access token is present": (r) =>
          r.json().data && r.json().data.access_token,
        "response status is 201": (r) => r.status === 201,
      });
    });

    group("Negative Registration - Duplicate Email", () => {
      const duplicateUser = generateTestUser();
      const registrationPayload = JSON.stringify(duplicateUser);
      const registrationParams = {
        headers: {
          "Content-Type": "application/json",
        },
      };

      // First attempt
      http.post(`${baseUrl}/register`, registrationPayload, registrationParams);

      // Second registration attempt with same email
      const duplicateRegistrationResponse = http.post(
        `${baseUrl}/register`,
        registrationPayload,
        registrationParams
      );

      check(duplicateRegistrationResponse, {
        "duplicate registration returns error": (r) =>
          r.json().status === "error",
        "error message is correct": (r) =>
          r.json().error === "Email already registered.",
        "response status is appropriate": (r) => r.status === 400,
      });
    });

    group("Negative Registration - Invalid Payloads", () => {
      const testCases = [
        { ...generateTestUser(), email: "" }, // Missing email
        { ...generateTestUser(), email: "invalid-email" }, // Invalid email format
        { ...generateTestUser(), password: "123" }, // Short password
        { ...generateTestUser(), name: "" }, // Missing name
      ];

      testCases.forEach((invalidPayload) => {
        const registrationPayload = JSON.stringify(invalidPayload);

        const registrationParams = {
          headers: {
            "Content-Type": "application/json",
          },
        };

        const invalidRegistrationResponse = http.post(
          `${baseUrl}/register`,
          registrationPayload,
          registrationParams
        );

        check(invalidRegistrationResponse, {
          "invalid payload returns error": (r) => r.json().status === "error",
          "response status is 400": (r) => r.status === 400,
        });
      });
    });
  });

  return validUser;
}
