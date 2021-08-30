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


const ids = {

    user_input_0 : "i_id0",
    user_input_1 : "i_id1",
    user_pw_input : "i_id2",
    login_form : "login_form",
    register_button : "register_button",
    login_button : "login_button",
    register_form : "register_form",
    pw_forget : "password_forget"

}



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
        const d0 = document.createElement("div");
        d0.classList.add("col-8");
        d0.id = ids.login_form;
        const d1 = document.createElement("input");
        d1.classList.add("col-12", "pt-0", "pb-0");
        d1.placeholder = "email";
        d1.type = "text";
        d1.id = ids.user_input_0
        //d1.onkeyup = () => { _internal.checkLoginEmail(undefined); }
        const d2 = document.createElement("input");
        d2.classList.add("col-12", "pt-0", "pb-0", "mt-2");
        d2.placeholder = "password";
        d2.type = "password";
        d2.id = ids.user_pw_input
        const d3 = document.createElement("button");
        d3.classList.add("col-12", "btn", "ros-bg-sb", "mt-2", "pt-0", "pb-0", "ml-0", "mr-0");
        d3.style.boxShadow = "1px 1px 6px #000;";
        d3.onmouseenter = () => {btn.onEnter(d3);}
        d3.onmouseleave = () => {btn.onLeave(d3); }
        d3.onclick = () => {_internal.login(); }
        d3.innerText = "Login";
        d3.id = ids.login_button

        const d4 = document.createElement("div")
        d4.classList.add("col-12", "text-right")
        console.log(http.p.password_forget)
        d4.id = ids.pw_forget
        d4.innerHTML = `<a class="ml-0 ros-pr text-right" style="font-size: 10px;" href="`+http.p.password_forget+`">Forgot Password?</a>`


        if (/Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent)){
            d0.classList.add("container-fluid")
            d0.classList.add("ros-bg-lb");
            d1.style.fontSize = '5em';
            d2.style.fontSize = '5em';
            d3.style.fontSize = '5em';
            d4.classList.add("mt-2")
            d4.innerHTML = `<a class="ml-0 ros-pr text-right mt-2" style="font-size: 25px;" href="`+http.p.password_forget+`">Forgot Password?</a>`
        }
        d0.appendChild(d1);
        d0.appendChild(d2);
        d0.appendChild(d3);
        d0.appendChild(d4);
        document.getElementById("login_form_container").appendChild(d0);
    },
    /*
    dynamically creates a register form for the frontend website
    */
    createRegisterForm : () => {
        const d0 = document.createElement("div");
        const d4 = document.createElement("button");
        const d1 = document.createElement("input");
        d0.classList.add("col-8");
        d0.id = ids.register_form;

        d4.classList.add("col-12", "btn", "ros-bg-sb", "mt-2", "pt-0", "pb-0", "ml-0", "mr-0");
        d4.style.boxShadow = "1px 1px 6px #000;";
        d4.onmouseenter = () => {
            btn.onEnter(d4);
        }
        d4.onmouseleave = () => {
            btn.onLeave(d4);
        }
        d4.id = ids.register_button;
        d4.innerText = "Register";

        d4.onclick = () => {
            _internal.register();
        }
        d1.classList.add("col-12", "pt-0", "pb-0");
        d1.placeholder = "email";
        d1.type ="text";
        d1.id = ids.user_input_0
        //d1.onkeyup = () => {_internal.checkRegisterEmail(undefined);};
        if (/Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent)) {
            d1.style.fontSize = '5em';
            d4.style.fontSize = '5em';
        }
        d0.appendChild(d1);
        d0.appendChild(d4);
        document.getElementById("login_form_container").appendChild(d0);
    },
    /*
        removes a created register form if it exists
     */
    removeRegisterForm : () =>  {
        var d0 = document.getElementById(ids.register_form);
        var d1 = document.getElementById(ids.user_input_0);
        var d4 = document.getElementById(ids.register_button);
        if (d0!=null || d0 != undefined)d0.remove();
        if (d1!=null || d1 != undefined)d1.remove();
        if (d4!=null || d4 != undefined)d4.remove();
    },
    /*
        removes a created login form if it exists
     */
    removeLoginForm : () => {
        var d0 = document.getElementById(ids.login_form);
        var d1 = document.getElementById(ids.user_input_0);
        var d2 = document.getElementById(ids.user_pw_input);
        var d3 = document.getElementById(ids.login_button);
        var d4 = document.getElementById(ids.pw_forget)
        if (d0!=null || d0 != undefined)d0.remove();
        if (d1!=null || d1 != undefined)d1.remove();
        if (d2!=null || d2 != undefined)d2.remove();
        if (d3!=null || d3 != undefined)d3.remove();
        if (d4!=null || d4 != undefined)d4.remove();
    },

    /*
        checks wether the username is already registered in the database
     */
    checkRegisterEmail : async function (cb_func) {
        var d = document.getElementById(ids.user_input_0);
        const r = new XMLHttpRequest();
        r.open(http.me.check, http.p.login + "?" + http.q.user_email + "=" + d.value , true);
        r.send();
        r.onreadystatechange = () => {
            if (r.readyState == XMLHttpRequest.DONE){
                var j = JSON.parse(r.response);
                console.log(j)
                if (!j[http.q.user_exist]){
                    d.style.color = "#A0CF00";
                    if (cb_func != undefined)cb_func(true)
                    const d4 = document.getElementById(ids.register_button);
                    d4.removeAttribute("disabled", "true")
                    document.getElementById(ids.pw_forget).remove()




                    return true
                }
                else{
                    d.style.color = color_sceme.dr;
                    if (cb_func != undefined)cb_func(false)
                    const d4 = document.getElementById(ids.register_button);
                    if (document.getElementById(ids.pw_forget)==null) {
                        const e0 = document.createElement("div")
                        e0.classList.add("col-12", "text-right")
                        e0.id = ids.pw_forget
                        e0.innerHTML = `<a class="ml-0 ros-pr text-right" style="font-size: 10px;" href="` + http.p.password_forget + `">Forgot Password?</a>`
                        document.getElementById(ids.register_form).appendChild(e0)
                    }

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

    checkLoginEmail : (cb_func) => {
        var d = document.getElementById(ids.user_input_0);
        const r = new XMLHttpRequest();
        r.open(http.me.check, http.p.login + "?" + http.q.user_email + "=" + d.value , true);
        r.send();
        r.onreadystatechange = () => {
            if (r.readyState == XMLHttpRequest.DONE){
                var j = JSON.parse(r.response);
                if (j[http.q.user_exist]){
                    if (cb_func!= undefined || cb_func!= null){
                        cb_func();
                    }
                    return true
                }
                else{
                    $(ids.login_button).style.backgroundColor = color_sceme.pr
                    setTimeout(() => {$(ids.login_button).style.backgroundColor = color_sceme.lb},500)
                    return false
                }
            }
        }


    },
    register : () => {
            _internal.checkRegisterEmail((r) => {
                if (r){

                    console.log("doing the registration")
                    const d1 = document.getElementById(ids.user_input_0);
                    console.log(d1)
                    if (d1 != null){
                        const email = d1.value
                        if (!email.includes("@")){
                            const d4 = document.getElementById(ids.register_button);
                            d4.classList.remove("ros-bg-sb")
                            d4.classList.add("ros-bg-dr");

                            const d0 = document.createElement("div")
                            d0.classList.add("text-center", "col-12")
                            if ( !email.includes("th-koeln"))d0.innerText = "The entered Email is not from TH Koeln"
                            if ( !email.includes("@"))d0.innerText = "The entered Text is not an Email"
                            d0.id = "err0"
                            d0.style.color = color_sceme.pr
                            document.getElementById(ids.register_form).appendChild(d0)
                            setTimeout(() => {
                                d4.classList.add("ros-bg-sb")
                                d4.classList.remove("ros-bg-dr");
                                const d1 = document.getElementById(ids.user_input_0);
                                d0.remove()
                            }, 3000);
                            return

                        }
                        const r = new XMLHttpRequest();
                        var e = http.p.register + "?" + http.q.user_email + "=" + email
                        r.open(http.me.post, e , true);
                        r.send();
                        r.onreadystatechange = () => {
                            if (r.readyState == XMLHttpRequest.DONE){
                                const d = JSON.parse(r.response);
                                if (d[http.q.success]){
                                    loadLoginForm();
                                }
                                else {
                                    const d4 = document.getElementById(ids.register_button);
                                    d4.classList.remove("ros-bg-sb")
                                    d4.classList.add("ros-bg-dr");
                                    setTimeout(() => {
                                        d4.classList.add("ros-bg-sb")
                                        d4.classList.remove("ros-bg-dr");
                                        const d1 = document.getElementById(ids.user_input_0);
                                        d1.value = "";
                                    }, 1000);

                                }
                            }
                        }
                    }
                }
            });
    },

    login : () => {
        _internal.checkLoginEmail(() => {
            const d1 = document.getElementById(ids.user_input_0);
            const d2 = document.getElementById(ids.user_pw_input);
            if (d1 != null && d2 != null) {
                const email = d1.value;
                const usp = encodeRot13(d2.value);
                const r = new XMLHttpRequest();
                r.open(http.me.post, http.p.login + "?" + http.q.user_email + "=" + email + "&" + http.q.user_password + "=" + usp, true);
                r.send();
                r.onreadystatechange = () => {
                    if (r.readyState == XMLHttpRequest.DONE) {
                        const d = JSON.parse(r.response);
                        console.log(d)
                        if (d[http.q.authenticated]){
                            $(ids.login_button).style.backgroundColor = "#A0CF00"
                            setTimeout(() => {$(ids.login_button).style.backgroundColor = color_sceme.lb
                                }
                                ,500)
                            setTimeout(() => {
                                window.open("http://"+ http.a.address + ":"+ http.a.port + "/" , "_self")
                                }
                                ,800)
                            return
                        }
                    else{
                            $(ids.login_button).style.backgroundColor = color_sceme.pr
                            setTimeout(() => {$(ids.login_button).style.backgroundColor = color_sceme.lb
                                }
                                ,500)
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

