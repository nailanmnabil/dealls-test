import { sleep } from "k6";
import { TestRegistration } from "./registerTest.js";
import { TestLogin } from "./loginTest.js";
import { TestSwipe } from "./swipeTest.js";

export const options = {
  vus: 1,
  iterations: 1,
  duration: "30s",
  thresholds: {
    http_req_duration: ["p(95)<500"],
  },
};

export default function () {
  const user = TestRegistration();
  sleep(1);
  const accessToken = TestLogin(user);
  sleep(1);
  TestSwipe(accessToken);
}
