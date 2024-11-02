import http from "k6/http";
import { check, group } from "k6";

export function TestSwipe(accessToken) {
  const baseUrl = "http://localhost:8080";

  let lastProfile;
  let premiumPackage;

  group("Swipe Profile Scenarios", function () {
    group("Positive Profile - Get Random Profile", () => {
      const getProfileParams = {
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${accessToken}`,
        },
      };

      const getProfileResponse = http.get(
        `${baseUrl}/profiles/random`,
        getProfileParams
      );
      check(getProfileResponse, {
        "get profile status is success": (r) => r.json().status === "success",
        "profile data is present": (r) =>
          r.json().data && r.json().data.user_id,
        "response status is 200": (r) => r.status === 200,
      });
    });

    group("Positive Profile - Get Random Profile Twice", () => {
      for (let i = 0; i < 2; i++) {
        const getProfileParams = {
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${accessToken}`,
          },
        };

        const getProfileResponse = http.get(
          `${baseUrl}/profiles/random`,
          getProfileParams
        );
        check(getProfileResponse, {
          "get profile status is success": (r) => r.json().status === "success",
          "profile data is present": (r) =>
            r.json().data && r.json().data.user_id,
          "profile data is same as before": (r) =>
            lastProfile ? r.json().data.user_id === lastProfile.user_id : true,
          "response status is 200": (r) => r.status === 200,
        });
        lastProfile = getProfileResponse.json().data;
      }
    });

    group("Negative Swipe - Invalid Profile Visit ID", () => {
      const swipeParams = {
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${accessToken}`,
        },
      };
      const swipePayload = JSON.stringify({
        profile_visit_id: "wrong-profile_visit_id",
        swipe_type: "LEFT",
      });
      const swipeResponse = http.post(
        `${baseUrl}/profiles/swipe`,
        swipePayload,
        swipeParams
      );
      check(swipeResponse, {
        "swipe status is error": (r) => r.json().status === "error",
        "response status is 400": (r) => r.status === 400,
      });
    });

    group("Positive Swipe - Swipe Profile", () => {
      const swipeParams = {
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${accessToken}`,
        },
      };
      const swipePayload = JSON.stringify({
        profile_visit_id: lastProfile.profile_visit_id,
        swipe_type: "LEFT",
      });
      const swipeResponse = http.post(
        `${baseUrl}/profiles/swipe`,
        swipePayload,
        swipeParams
      );
      check(swipeResponse, {
        "swipe status is success": (r) => r.json().status === "success",
        "response status is 200": (r) => r.status === 200,
      });
    });

    group("Negative Swipe - Swipe Twice", () => {
      const swipeParams = {
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${accessToken}`,
        },
      };
      const swipePayload = JSON.stringify({
        profile_visit_id: lastProfile.profile_visit_id,
        swipe_type: "LEFT",
      });
      const swipeResponse = http.post(
        `${baseUrl}/profiles/swipe`,
        swipePayload,
        swipeParams
      );
      check(swipeResponse, {
        "swipe status is error": (r) => r.json().status === "error",
        "swipe error message is correct": (r) =>
          r.json().error === "Profile already swiped.",
        "response status is 400": (r) => r.status === 400,
      });
    });

    group("Positive Swipe - Swipe 10 Times", () => {
      for (let i = 0; i < 9; i++) {
        // 9 because we already use 1 quota on above test case
        const getProfileParams = {
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${accessToken}`,
          },
        };
        const getProfileResponse = http.get(
          `${baseUrl}/profiles/random`,
          getProfileParams
        );
        check(getProfileResponse, {
          "get profile status is success": (r) => r.json().status === "success",
          "profile data is present": (r) =>
            r.json().data && r.json().data.user_id,
          "response status is 200": (r) => r.status === 200,
        });

        const swipeParams = {
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${accessToken}`,
          },
        };
        const swipePayloads = JSON.stringify({
          profile_visit_id: getProfileResponse.json().data.profile_visit_id,
          swipe_type: "LEFT",
        });
        const swipeResponse = http.post(
          `${baseUrl}/profiles/swipe`,
          swipePayloads,
          swipeParams
        );
        check(swipeResponse, {
          "swipe status is success": (r) => r.json().status === "success",
          "response status is 200": (r) => r.status === 200,
        });
      }
    });

    group("Negative Profile - Get Random Profile While No Left Quota", () => {
      const getProfileParams = {
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${accessToken}`,
        },
      };

      const getProfileResponse = http.get(
        `${baseUrl}/profiles/random`,
        getProfileParams
      );
      check(getProfileResponse, {
        "get profile status is error": (r) => r.json().status === "error",
        "response status is 403": (r) => r.status === 403,
      });
    });

    group("Positive Purchase - Get All Premium Packages", () => {
      const getPremiumPackages = {
        headers: {
          "Content-Type": "application/json",
        },
      };

      const getPackagesResponse = http.get(
        `${baseUrl}/packages`,
        getPremiumPackages
      );
      check(getPackagesResponse, {
        "get all packages is success": (r) => r.json().status === "success",
        "get all packages data is exist": (r) => r.json().data,
        "response status is 200": (r) => r.status === 200,
      });
      getPackagesResponse.json().data.forEach((pkg) => {
        if (pkg.feature_type === "UNLIMITED_SWIPE_QUOTA") {
          premiumPackage = pkg;
        }
      });
    });

    group("Negative Purchase - Buy With Wrong Package ID", () => {
      const buyParams = {
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${accessToken}`,
        },
      };
      const buyPayload = JSON.stringify({
        package_id: "wrong-package-id",
      });
      const buyResponse = http.post(
        `${baseUrl}/profiles/swipe`,
        buyPayload,
        buyParams
      );
      check(buyResponse, {
        "buy package status is error": (r) => r.json().status === "error",
        "response status is 400": (r) => r.status === 400,
      });
    });

    group("Positive Purchase - Buy Package", () => {
      const buyParams = {
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${accessToken}`,
        },
      };
      const buyPayload = JSON.stringify({
        package_id: premiumPackage.package_id,
      });
      const buyResponse = http.post(
        `${baseUrl}/packages/purchase`,
        buyPayload,
        buyParams
      );
      check(buyResponse, {
        "buy package status is success": (r) => r.json().status === "success",
        "response status is 200": (r) => r.status === 200,
      });
    });

    group(
      "Positive Profile - Get Random Profile With UNLIMITED_SWIPE_QUOTA Package",
      () => {
        const getProfileParams = {
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${accessToken}`,
          },
        };

        const getProfileResponse = http.get(
          `${baseUrl}/profiles/random`,
          getProfileParams
        );
        check(getProfileResponse, {
          "get profile status is success": (r) => r.json().status === "success",
          "response status is 200": (r) => r.status === 200,
        });
      }
    );
  });
}
