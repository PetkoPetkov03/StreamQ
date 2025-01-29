interface UserImpl {
    build: () => boolean
}

type User = UserImpl & {
    email: string,
    username: string,
    password: string,
    cpassword: string,
}

const fetchRegisterFunc = async(user: User) => {
    const response = await fetch("/register", {
        method: "POST",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify(user)
    })

    if (response.ok != true) {
        console.error((await response.json()).err);
        return;
    }

    location.replace("/");
}

const callRegisterProcess = (event: FormDataEvent) => {
    event.preventDefault();

    const user: User = {
        email: getInnerHTML("email"),
        username: getInnerHTML("username"),
        password: getInnerHTML("password"),
        cpassword: getInnerHTML("cpassword"),
        build: () => {
            if(user.email.length == 0 || user.password.length == 0 || user.cpassword.length == 0 || user.username.length == 0) {
                return false;
            }
            return true;
        },
    }

    if(!user.build()) {
        console.log(user);
        
        throw new Error("User construct failure on front end!");
    }

    fetchRegisterFunc(user);
}