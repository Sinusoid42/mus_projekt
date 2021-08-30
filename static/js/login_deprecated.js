/**
 MIT Licence:

 Copyright 2021 (c)ROS_Project_Group {TH Koeln - 6.Semester - Medientechnologie}

 Permission is hereby granted, free of charge, to any person obtaining a copy of this
 software and associated documentation files (the "Software"), to deal in the Software
 without restriction, including without limitation the rights to use, copy, modify, merge,
 publish, distribute, sublicense, and/or sell copies of the Software, and to permit
 persons to whom the Software is furnished to do so, subject to the following conditions:

 The above copyright notice and this permission notice shall be included in all copies or
 substantial portions of the Software.

 THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED,
 INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR
 PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE
 FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
 OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
 DEALINGS IN THE SOFTWARE.


 - - - - - - - - - - - - - - - - - - - - -

 Main JS component impplementation
 to accomodate for all functionalities the frontend is required to have in this project

 This file encapsules all functionalities in const placeholder objects for conveniance

 @author benwinter (06.07.2021)

 */


window.onload = () => {
    api.resolveServerAddress.then((server_data_json) => {
    console.log(server_data_json)
        http.putServerData(server_data_json);
        loadLoginForm();
    })
}
const btn = {

    onEnter : (e) => {
        e.classList.remove('ros-bg-sb');
        e.classList.add('ros-bg-lb');
        e.style.boxShadow = '1px 1px 10px #CEE3F2';
    },

    onLeave : (e) => {
        e.classList.remove('ros-bg-lb');
        e.classList.add('ros-bg-sb');
        e.style.boxShadow = '1px 1px 6px #273A4C';
    }
}

const _internal = {
    /*
        dynamically creates a login form for the frontend website
    */
    createLoginForm : () => {


        if (/Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent)){

            const d0 = document.createElement("div");
            d0.classList.add("col-8");
            d0.classList.add("container-fluid")
            d0.id = "login_form";
            d0.classList.add("ros-bg-lb");

            const d1 = document.createElement("input");
            d1.classList.add("col-12", "pt-0", "pb-0");
            d1.placeholder = "username";
            d1.type = "text";
            d1.style.fontSize = '5em';
            d1.id = "user_name_login";
            d1.onkeyup = () => {
                _internal.checkLoginUsername(undefined);
            }

            const d2 = document.createElement("input");
            d2.classList.add("col-12", "pt-0", "pb-0", "mt-2");
            d2.placeholder = "password";
            d2.type = "password";
            d2.id = "user_name_password";
            d2.style.fontSize = '5em';
            const d3 = document.createElement("button");
            d3.classList.add("col-12", "btn", "ros-bg-sb", "mt-2", "pt-0", "pb-0", "ml-0", "mr-0");
            d3.style.boxShadow = "1px 1px 6px #000;";
            d3.onmouseenter = () => {
                btn.onEnter(d3);
            }
            d3.style.fontSize = '5em';
            d3.onmouseleave = () => {
                btn.onLeave(d3);
            }
            d3.onclick = () => {
                _internal.login();
            }
            d3.innerText = "Login";
            d3.id = "login_button";

            d0.appendChild(d1);
            d0.appendChild(d2);
            d0.appendChild(d3);
            document.getElementById("login_form_container").appendChild(d0);


        }
        else {


            const d0 = document.createElement("div");
            d0.classList.add("col-8");
            d0.id = "login_form";

            const d1 = document.createElement("input");
            d1.classList.add("col-12", "pt-0", "pb-0");
            d1.placeholder = "username";
            d1.type = "text";
            d1.id = "user_name_login";
            d1.onkeyup = () => {
                _internal.checkLoginUsername(undefined);
            }

            const d2 = document.createElement("input");
            d2.classList.add("col-12", "pt-0", "pb-0", "mt-2");
            d2.placeholder = "password";
            d2.type = "password";
            d2.id = "user_name_password";

            const d3 = document.createElement("button");
            d3.classList.add("col-12", "btn", "ros-bg-sb", "mt-2", "pt-0", "pb-0", "ml-0", "mr-0");
            d3.style.boxShadow = "1px 1px 6px #000;";
            d3.onmouseenter = () => {
                btn.onEnter(d3);
            }
            d3.onmouseleave = () => {
                btn.onLeave(d3);
            }
            d3.onclick = () => {
                _internal.login();
            }
            d3.innerText = "Login";
            d3.id = "login_button";

            d0.appendChild(d1);
            d0.appendChild(d2);
            d0.appendChild(d3);
            document.getElementById("login_form_container").appendChild(d0);
        }
    },
    /*
    dynamically creates a register form for the frontend website
    */
    createRegisterForm : () => {

        if (/Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent)) {
            const d0 = document.createElement("div");
            d0.classList.add("col-8");
            d0.id = "register_form";

            const d1 = document.createElement("input");
            d1.style.fontSize = '5em';
            d1.classList.add("col-12", "pt-0", "pb-0");
            d1.placeholder = "username";
            d1.type ="text";
            d1.id ="user_name_login";
            d1.onkeyup = () => {
                _internal.checkRegisterUsername(undefined);
            }
            const d2 = document.createElement("input");
            d2.style.fontSize = '5em';
            d2.classList.add("col-12", "pt-0", "pb-0", "mt-2");
            d2.placeholder = "password";
            d2.type ="password";
            d2.id ="user_name_password_0";
            d2.onkeyup = () => {
                _internal.checkPasswords();
            }

            const d3 = document.createElement("input");
            d3.style.fontSize = '5em';
            d3.classList.add("col-12", "pt-0", "pb-0", "mt-2");
            d3.placeholder = "password";
            d3.type ="password";
            d3.id ="user_name_password_1";
            d3.onkeyup = () => {
                _internal.checkPasswords();
            }

            const d4 = document.createElement("button");
            d4.style.fontSize = '5em';
            d4.classList.add("col-12", "btn", "ros-bg-sb", "mt-2", "pt-0", "pb-0", "ml-0", "mr-0");
            d4.style.boxShadow = "1px 1px 6px #000;";
            d4.onmouseenter = () => {
                btn.onEnter(d4);
            }
            d4.onmouseleave = () => {
                btn.onLeave(d4);
            }
            d4.id = "register_button";
            d4.innerText = "Register";
            d0.appendChild(d1);
            d0.appendChild(d2);
            d0.appendChild(d3);
            d0.appendChild(d4);
            d4.onclick = () => {
                _internal.register();
            }
            document.getElementById("login_form_container").appendChild(d0);

        }
        else {
            const d0 = document.createElement("div");
            d0.classList.add("col-8");
            d0.id = "register_form";

            const d1 = document.createElement("input");
            d1.classList.add("col-12", "pt-0", "pb-0");
            d1.placeholder = "username";
            d1.type ="text";
            d1.id ="user_name_login";
            d1.onkeyup = () => {
                _internal.checkRegisterUsername(undefined);
            }
            const d2 = document.createElement("input");
            d2.classList.add("col-12", "pt-0", "pb-0", "mt-2");
            d2.placeholder = "password";
            d2.type ="password";
            d2.id ="user_name_password_0";
            d2.onkeyup = () => {
                _internal.checkPasswords();
            }

            const d3 = document.createElement("input");
            d3.classList.add("col-12", "pt-0", "pb-0", "mt-2");
            d3.placeholder = "password";
            d3.type ="password";
            d3.id ="user_name_password_1";
            d3.onkeyup = () => {
                _internal.checkPasswords();
            }

            const d4 = document.createElement("button");
            d4.classList.add("col-12", "btn", "ros-bg-sb", "mt-2", "pt-0", "pb-0", "ml-0", "mr-0");
            d4.style.boxShadow = "1px 1px 6px #000;";
            d4.onmouseenter = () => {
                btn.onEnter(d4);
            }
            d4.onmouseleave = () => {
                btn.onLeave(d4);
            }
            d4.id = "register_button";
            d4.innerText = "Register";
            d0.appendChild(d1);
            d0.appendChild(d2);
            d0.appendChild(d3);
            d0.appendChild(d4);
            d4.onclick = () => {
                _internal.register();
            }
            document.getElementById("login_form_container").appendChild(d0);


        }





    },
    /*
        removes a created register form if it exists
     */
     removeRegisterForm : () =>  {
        var d0 = document.getElementById("register_form");
        var d1 = document.getElementById("user_name_login");
        var d2 = document.getElementById("user_name_password_0");
        var d3 = document.getElementById("user_name_password_1");
        var d4 = document.getElementById("register_button");
        if (d0!=null || d0 != undefined)d0.remove();
        if (d1!=null || d1 != undefined)d1.remove();
        if (d2!=null || d2 != undefined)d2.remove();
        if (d3!=null || d3 != undefined)d3.remove();
        if (d4!=null || d4 != undefined)d4.remove();
    },
    /*
        removes a created login form if it exists
     */
    removeLoginForm : () => {
        var d0 = document.getElementById("login_form");
        var d1 = document.getElementById("user_name_login");
        var d2 = document.getElementById("user_name_password_0");
        var d3 = document.getElementById("login_button");
        if (d0!=null || d0 != undefined)d0.remove();
        if (d1!=null || d1 != undefined)d1.remove();
        if (d2!=null || d2 != undefined)d2.remove();
        if (d3!=null || d3 != undefined)d3.remove();
    },

    /*
        checks the input fields if both passwords have to be entered
     */
    checkPasswords : () => {
        var d0 = document.getElementById("user_name_password_0");
        var d1 = document.getElementById("user_name_password_1");
        if (d0 != null && d1 != null){

            if (d0.value.length ==0 || d1.value.length == 0){
                d0.style.color = "#000";
                d1.style.color = "#000";
                return false
            } else if (d0.value == d1.value && d0.value.length > 6 && d1.value.length > 6) {
                d0.style.color = "#A0CF00";
                d1.style.color = "#A0CF00";
                console.log("password check succeeded")
                return true
            }
            else{
                d0.style.color = "#FF4040";
                d1.style.color = "#FF4040";
                return false
            }
        }
        return false
    },

    /*
        checks wether the username is already registered in the database
     */
    checkRegisterUsername : async function (cb_func) {
        var d = document.getElementById("user_name_login");
        const r = new XMLHttpRequest();
        r.open(http.me.check, "/" + http.p.login + "?" + http.q.user_name + "=" + d.value , true);
        r.send();
        r.onreadystatechange = () => {
            if (r.readyState == XMLHttpRequest.DONE){
                var j = JSON.parse(r.response);
                if (!j[http.q.user_exist]){
                    d.style.color = "#A0CF00";
                    if (cb_func != undefined)cb_func(true)
                    const d4 = document.getElementById("register_button");
                    d4.removeAttribute("disabled", "true")
                    return true
                }
                else{
                    d.style.color = color_sceme.dr;
                    if (cb_func != undefined)cb_func(false)
                    const d4 = document.getElementById("register_button");
                    d4.classList.remove("ros-bg-sb")
                    d4.classList.add("ros-bg-dr");
                    d4.setAttribute("disabled", "true")
                    setTimeout(() => {
                        d4.classList.add("ros-bg-sb")
                        d4.classList.remove("ros-bg-dr");
                    }, 1000)
                    return false
                }
            }
        }
    },
    checkLoginUsername : (cb_func) => {

        var d = document.getElementById("user_name_login");
        const r = new XMLHttpRequest();
        r.open(http.me.check, "/" + http.p.login + "?" + http.q.user_name + "=" + d.value , true);
        r.send();
        r.onreadystatechange = () => {
            if (r.readyState == XMLHttpRequest.DONE){
                var j = JSON.parse(r.response);
                if (j[http.q.user_exist]){
                    const d4 = document.getElementById("login_button");
                    d4.removeAttribute("disabled", "true")
                    if (cb_func!= undefined || cb_func!= null){
                        cb_func();
                    }
                    return true
                }
                else{

                    const d4 = document.getElementById("login_button");
                    d4.setAttribute("disabled", "true")
                    return false
                }
            }
        }


    },
    register : () => {
        if (_internal.checkPasswords()){
            _internal.checkRegisterUsername((r) => {



                if (r){
                    const d1 = document.getElementById("user_name_login");
                    const d2 = document.getElementById("user_name_password_0");
                    if (d1 != null && d2 != null){
                        const usn = d1.value
                        const usp = encodeRot13(d2.value)
                        const r = new XMLHttpRequest();
                        r.open(http.me.post, http.p.register + "?" + http.q.user_name + "=" + usn + "&" + http.q.user_password +"=" + usp, true);
                        r.send();
                        r.onreadystatechange = () => {
                            if (r.readyState == XMLHttpRequest.DONE){
                                const d = JSON.parse(r.response);
                                if (d[http.q.success]){
                                    loadLoginForm();
                                }
                                else {
                                    const d4 = document.getElementById("register_button");
                                    d4.classList.remove("ros-bg-sb")
                                    d4.classList.add("ros-bg-dr");
                                    setTimeout(() => {
                                        d4.classList.add("ros-bg-sb")
                                        d4.classList.remove("ros-bg-dr");
                                        const d1 = document.getElementById("user_name_login");
                                        d1.value = "";
                                    }, 1000);

                                }
                            }
                        }
                    }
                }
            });
        }
    },

    login : () => {
        _internal.checkLoginUsername(() => {
            const d1 = document.getElementById("user_name_login");
            const d2 = document.getElementById("user_name_password");
            if (d1 != null && d2 != null) {
                const usn = d1.value;
                const usp = encodeRot13(d2.value);
                const r = new XMLHttpRequest();
                r.open(http.me.post, http.p.login + "?" + http.q.user_name + "=" + usn + "&" + http.q.user_password + "=" + usp, true);
                r.send();
                r.onreadystatechange = () => {
                    if (r.readyState == XMLHttpRequest.DONE) {
                        const d = JSON.parse(r.response);
                        if (d[http.q.authenticated]){
                            console.log("The user will now be authenticated")
                            window.open("http://"+ http.a.address + ":"+ http.a.port + "/" , "_self")


                        }
                    }
                }
            }
        });
    }
}







/*
    loads the login form
 */
const loadLoginForm = () => {
    _internal.removeRegisterForm();
    _internal.removeLoginForm();
    _internal.createLoginForm();
}
/*
    loads the register form
 */
const loadRegisterForm = () => {
    _internal.removeRegisterForm();
    _internal.removeLoginForm();
    _internal.createRegisterForm();
}

