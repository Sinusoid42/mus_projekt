/*
    password reset client implementations to send reset data
    this implementation focusses on the commuication between client and server
    in order to create a reset REQUEST for a user account password
 */
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

 Templates api website

 this frontend webpage delivers functionalities to create a new password reset request
 if
 */


const pwr = {
    check : (cb_func, err) => {
        var d = document.getElementById("i0");
        const r = new XMLHttpRequest();
        r.open(http.me.post, http.p.password_forget + "?" + http.q.user_email + "=" + d.value, true);
        r.send();
        r.onreadystatechange = () => {
            if (r.readyState == XMLHttpRequest.DONE) {
                var j = JSON.parse(r.response);
                console.log(j)
                if (j[http.q.user_exist]) {
                    cb_func(j)
                } else {
                    err(j)
                }
            }
        }
    },
}
const a = ()=> {
    document.getElementById("a0")
    pwr.check((e)=> {
        if (e["answer0"]!=null)document.getElementById("a0").innerText = e["answer0"]
        if (e["answer1"]!=null)document.getElementById("a1").innerText = e["answer1"]
    }, (e) => {
        console.log("hello world")
        if (e["answer0"]!=null)document.getElementById("a0").innerText = e["answer0"]
        if (e["answer1"]!=null)document.getElementById("a1").innerText = e["answer1"]
        var a = "http://" + http.a.address + ":" + http.a.port + "/login"
        console.log(a)
        setTimeout(() => {window.open(a, "_self")},2500);
    })
}

window.onload = () => {
    api.resolveServerAddress.then((e) => {
        http.putServerData(e)
    });
}