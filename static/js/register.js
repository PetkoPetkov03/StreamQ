"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
const fetchRegisterFunc = (user) => __awaiter(void 0, void 0, void 0, function* () {
    const response = yield fetch("/register", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(user)
    });
    if (response.ok != true) {
        console.error((yield response.json()).err);
        return;
    }
    location.replace("/");
});
const callRegisterProcess = (event) => {
    event.preventDefault();
    const user = {
        email: getInnerHTML("email"),
        username: getInnerHTML("username"),
        password: getInnerHTML("password"),
        cpassword: getInnerHTML("cpassword"),
        build: () => {
            if (user.email.length == 0 || user.password.length == 0 || user.cpassword.length == 0 || user.username.length == 0) {
                return false;
            }
            return true;
        },
    };
    if (!user.build()) {
        console.log(user);
        throw new Error("User construct failure on front end!");
    }
    fetchRegisterFunc(user);
};
