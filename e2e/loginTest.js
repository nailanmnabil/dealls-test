import http from "k6/http";
import { check, group } from "k6";

export function TestLogin(user) {
  const baseUrl = "http://localhost:8080";
  const loginPayload = {
    email: user.email,
    password: user.password,
  };
  let accessToken;

  group("User Login Scenarios", function () {
    group("Positive Login", () => {
      const positiveLoginPayload = JSON.stringify(loginPayload);
      const loginParams = {
        headers: {
          "Content-Type": "application/json",
        },
      };

      const loginResponse = http.post(
        `${baseUrl}/login`,
        positiveLoginPayload,
        loginParams
      );
      check(loginResponse, {
        "login status is success": (r) => r.json().status === "success",
        "access token is present": (r) =>
          r.json().data && r.json().data.access_token,
        "response status is 200": (r) => r.status === 200,
      });

      accessToken = loginResponse.json().data.access_token;
    });

    group("Negative Login - Wrong Email Password", () => {
      const testCases = [
        { ...loginPayload, email: "wrong@example.com" },
        { ...loginPayload, password: "wrongwrongwrong" },
      ];

      testCases.forEach((invalidPayload) => {
        const loginPayload = JSON.stringify(invalidPayload);
        const loginParams = {
          headers: {
            "Content-Type": "application/json",
          },
        };

        const invalidLoginResponse = http.post(
          `${baseUrl}/login`,
          loginPayload,
          loginParams
        );

        check(invalidLoginResponse, {
          "invalid payload returns error": (r) => r.json().status === "error",
          "response status is 401": (r) => r.status === 401,
        });
      });
    });

    group("Negative Login - Invalid Payloads", () => {
      const testCases = [
        { ...loginPayload, email: "" }, // Missing email
        { ...loginPayload, email: "invalid-email" }, // Invalid email format
        { ...loginPayload, password: "" }, // Missing password
        { ...loginPayload, password: "123" }, // Short password
      ];

      testCases.forEach((invalidPayload) => {
        const loginPayload = JSON.stringify(invalidPayload);

        const loginParams = {
          headers: {
            "Content-Type": "application/json",
          },
        };

        const invalidLoginResponse = http.post(
          `${baseUrl}/login`,
          loginPayload,
          loginParams
        );

        check(invalidLoginResponse, {
          "invalid payload returns error": (r) => r.json().status === "error",
          "response status is 400": (r) => r.status === 400,
        });
      });
    });
  });

  return accessToken;
}
