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

 Arduino addition website

 Delivers functionality and responsiveness when attaching a new arduino to the system

 this Website deilvers the functionlities for checking on site the connections and infos
 about all connected arduinos for the server, possibility to add a new one and remove one
 as well as track last updates and infos


 @author ben
 */

let selected = false
let s_e = null


const settings = {

    db_id : "",
    db_rev : "",
    arduino_id : "",
    room_location : "",
    room_id : "",
    arduino_ip_address : "",
    arduino_inet_port : "",
    arduino_password : "",
    arduino_type : "",
    arduino_firmware : "",
    arduino_active_template : "",
    arduino_selected_id : "",


    push : () => {
        settings.arduino_selected_id = settings.arduino_id
        document.getElementById("microcontroller_settings_id").value = settings.arduino_id
        document.getElementById("microcontroller_settings_db_id").innerText = settings.db_id
        document.getElementById("microcontroller_settings_db_rev").innerText = settings.db_rev
        document.getElementById("microcontroller_settings_room").innerText = settings.room_location
        document.getElementById("microcontroller_settings_ip").value = settings.arduino_ip_address
        document.getElementById("microcontroller_settings_port").value = settings.arduino_inet_port
        document.getElementById("microcontroller_settings_password").value = settings.arduino_password
        document.getElementById("microcontroller_type").value = settings.arduino_type
        document.getElementById("microcontroller_firmware").value = settings.arduino_firmware
        document.getElementById("microcontroller_settings_current_template").value = settings.arduino_active_template
        document.getElementById("microcontroller_settings_room_id").value = settings.room_id
    },

    fetch : () => {
        settings.arduino_selected_id = settings.arduino_id
        settings.arduino_id = document.getElementById("microcontroller_settings_id").value
        settings.db_id = document.getElementById("microcontroller_settings_db_id").innerText
        settings.db_rev = document.getElementById("microcontroller_settings_db_rev").innerText
        settings.arduino_type = document.getElementById("microcontroller_type").value
        settings.arduino_firmware = document.getElementById("microcontroller_firmware").value
        settings.arduino_ip_address = document.getElementById("microcontroller_settings_ip").value
        settings.arduino_inet_port = document.getElementById("microcontroller_settings_port").value
        settings.arduino_password = document.getElementById("microcontroller_settings_password").value
        settings.arduino_active_template = document.getElementById("microcontroller_settings_current_template").value
        settings.room_id = document.getElementById("microcontroller_settings_room_id").value
    }
}

const cs = (title, content, id) => {
    return '<div class ="row justify-content-center text-center" id="' + id +'">' +
        '\n<b>' + title +'</b> ' + content + '\n</div>\n'
}

const select = function(e, id){
    if (s_e!=null){
        s_e.style.backgroundColor = color_sceme.sb
    }
    deselect()
    document.getElementById("microcontroller_settings_save_btn").removeAttribute("disabled", "true")
    document.getElementById("microcontroller_settings_delete_btn").removeAttribute("disabled", "true")
    selected = true
    s_e = e
    e.style.backgroundColor = color_sceme.pr

    const r = new XMLHttpRequest()
    r.open(http.me.select, http.p.arduino_api +"?microcontroller_id="+id, true)
    r.send()
    r.onreadystatechange = () => {
        if (r.readyState == XMLHttpRequest.DONE){
            e = JSON.parse(r.response)
            put_settings(e)
            settings.push()
        }
    }
}

const deselect = function(){
    if (s_e != null){
        s_e.style.backgroundColor = color_sceme.sb
    }
    const a = document.getElementById("deletion_form")
    if (a!=null)a.remove()
    document.getElementById("microcontroller_settings_save_btn").setAttribute("disabled", "true")
    document.getElementById("microcontroller_settings_delete_btn").setAttribute("disabled", "true")
    selected = false
    s_e = null
}

const  create_arduino_info_field = (jc) => {
    console.log(jc)
    const e0 = document.createElement("div")
    e0.id = jc["microcontroller_id"] + "_selection_field"
    e0.classList.add("row", "justify-content-center", "mt-2", "mb-2")
    e0.innerHTML = '<div class="col-10 ros-bg-sb" style="border-style: solid;' +
        'border-radius: 1em" onclick="select(this, `' +  jc["microcontroller_id"] + '`)" ' +
        'onmouseenter="onenter(this)"' +
        'onmouseleave="onleave(this)">'+
        cs("Room:", jc["room_location"], jc["microcontroller_id"] +"_room_location") +
        cs("Microcontroller ID:", jc["microcontroller_id"], jc["microcontroller_id"] +"_microcontroller_id") +
        cs("Room ID:", jc["room_id"], jc["microcontroller_id"] +"_room_id") +
        cs("Last Update:", jc["last_fetch"], jc["microcontroller_id"] +"_last_fetch") +
        cs("Last Info:", jc["last_info"], jc["microcontroller_id"] +"_last_info") +
        cs("Status: ", jc["status"], jc["microcontroller_id"] +"_status")+'</div>'
    document.getElementById("microcontroller_updates_container").appendChild(e0)
}

const onenter = (e) => {
    if (selected && e == s_e)return;
    e.style.backgroundColor = color_sceme.lr
}
const onleave = (e) => {
    if (selected && e == s_e)return;
    e.style.backgroundColor = color_sceme.sb
}

const fetch_arduino_data = () => {
    const r = new XMLHttpRequest()
    r.open(http.me.fetch, http.p.arduino_api, true)
    r.send()
    r.onreadystatechange = () => {
        if (r.readyState == XMLHttpRequest.DONE){
            const e = JSON.parse(r.response)
            e["data"].forEach((k) => {
                create_arduino_info_field(k)
            })
        }
    }
}

const put_settings = (e) => {
    settings.db_id = e["_id"]
    settings.db_rev = e["_rev"]
    settings.arduino_id = e["microcontroller_id"]
    settings.arduino_ip_address = e["microcontroller_ip_address"]
    settings.arduino_inet_port = e["microcontroller_inet_port"]
    settings.arduino_password = e["microcontroller_password"]
    settings.arduino_firmware = e["microcontroller_firmware"]
    settings.arduino_active_template = e["microcontroller_active_template"]
    settings.arduino_type = e["microcontroller_type"]
    settings.room_location = e["room_location"]
    settings.room_id = e["room_id"]
    settings.push()
}

const deleteSelectedArduino = () => {
    const e = document.getElementById("microcontroller_settings_container")
    const se0 = document.createElement("div")
    se0.id ="deletion_form";
    se0.classList.add("row", "justify-content-center", "mt-4", "mb-3")
    se0.innerHTML = `<div class="col-12"><div class="row justify-content-center">
    <div class="col-8">Enter the Arduino ID of the selected Arduino to confirm the deletion
    </div></div><div class="row justify-content-center mt-2"><div class="col-4"><input 
    class="input-group" onkeyup="checkDeletionIds(this.value)" id="confirmation_input"></div><div class="col-4"><button 
    class="btn" id="confirmation_btn" style="background-color: `+ color_sceme.lr +`" onclick="checkConfirmation()">Confirm Deletion</button></div></div></div>`
    e.appendChild(se0)
}

const checkConfirmation = () => {
    const a = document.getElementById("confirmation_input")
    if (a!=null && a.value == settings.arduino_id){
        //confirmation is fine
        //now handle server communication
        const r = new XMLHttpRequest()
        r.open(http.me.delete,http.p.arduino_api + "?microcontroller_id="+settings.arduino_id, true)
        r.send()
        r.onreadystatechange = () => {
            if (r.readyState == XMLHttpRequest.DONE){
                console.log(r.response)
                if (JSON.parse(r.response)["success"]){
                    console.log("deletion:success")
                    const b = document.getElementById(settings.arduino_id + "_selection_field");
                    if (b!=null)b.remove();
                    deselect();
                }
                else{
                    //TODO Catch error when deletion was unsuccessful
                }
            }
        }
    }
}

const checkDeletionIds = (a) => {
    if (a == settings.arduino_id){
        document.getElementById("confirmation_btn").style.backgroundColor = "#80F530"
    }else{
        document.getElementById("confirmation_btn").style.backgroundColor = color_sceme.lr
    }

}

const safeSelectedArduino = () => {
    settings.fetch()
    const r = new XMLHttpRequest()
    const query =
        "id=" +  settings.arduino_selected_id + "&" +
        "microcontroller_id=" + settings.arduino_id + "&" +
        "microcontroller_ip_address=" + settings.arduino_ip_address + "&" +
        "microcontroller_inet_port=" + settings.arduino_inet_port +"&"+
        "microcontroller_password=" + settings.arduino_password +"&"+
        "microcontroller_firmware=" + settings.arduino_firmware +"&"+
        "microcontroller_active_template=" + settings.arduino_active_template +"&"+
        "microcontroller_type=" + settings.arduino_type + "&" +
        "room_id=" + settings.room_id
    const addr = http.p.arduino_api + "?" + query
    console.log(addr)



    r.open(http.me.put, addr,  true)
    r.send()
    r.onreadystatechange = () => {
        if (r.readyState == XMLHttpRequest.DONE){
            const e = JSON.parse(r.response)
            if (e["success"]){
                settings.room_location = e["room_location"]
                settings.push()
                console.log(settings)

                swap(settings.arduino_selected_id, settings.arduino_id, "_room_location", '<b>Room: </b>' + settings.room_location)
                swap(settings.arduino_selected_id, settings.arduino_id, "_room_id", '<b>Room ID: </b>' + settings.room_id)
                swap(settings.arduino_selected_id, settings.arduino_id, "_microcontroller_id", '<b>Arduino ID: </b>' + settings.arduino_id)
                swap(settings.arduino_selected_id, settings.arduino_id, "_room_location", '<b>Arduino ID: </b>' + settings.room_location)

                deselect()
            }
        }
    }
}

const swap = (id_prev, id_update, ext, content) => {
    let a = document.getElementById(id_prev + ext)
    a.innerHTML = content
    a.id = id_update + ext
}

const create_new_arduino = () => {
    const r = new XMLHttpRequest()
    r.open(http.me.post, http.p.arduino_api, true)
    r.send()
    r.onreadystatechange = () => {
        if (r.readyState == XMLHttpRequest.DONE){
            const e = JSON.parse(r.response)
            create_arduino_info_field(e);
            put_settings(e)
        }
    }
}

window.onload = () => {
    api.resolveServerAddress.then(fetch_arduino_data())
}


