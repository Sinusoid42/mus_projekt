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

 Rooms addition website

 This api file is contained within namespaces for better scoping of variables
 The RAPI Interface delivers functionalities in order to manipulate the dom-tree
 corresponding to the needs the website

 rooms can be dynamically seleted and removed, the server will respond async to those
 requests and no reloading of databases has to be done

 Including Variable scopes the "settings" tag is responsible for upholding all
 catched variables necessary to run this frontend application


 @author ben
 */

const settings = {

    db_id : "",
    db_rev : "",
    arduino_attached : false,
    arduino_id : "",
    room_bookable: false,
    room_name : "",                 //eg: Hörsaal1, Praktikums Labor EM1
    room_name_misc : "",             //eg: ChillerEcke, Höhle des Löwen
    room_id : "",
    max_people : 0,
    max_booking_duration : 0,
    room_location : {
        floor : 0,
        wing : "",
        nbr : 0,
        name : "",
    },
}


const rapi = {
    btn : "rbtn",
    sf : 0,         //selected floor
    sw : "N",       //selected wing
    w : null,       //the wing element
    f : null,       //the floor element
    r : null,       //the room element
    fc : "floor_container",
    wc : "wing_container",
    rc : "room_container",
    floors : 13,

    settings_ids : "rs_",

    room_id : "rid",
    room_name : "rnm",
    room_name_misc : "rmn",
    room_arduino_available: "raa",
    room_bookable: "rbb",
    room_max_capacity: "rmx",
    room_max_booking_duration : "mbd",

    location_room_name : "lrnm",
    location_room_floor_level: "lrf",
    location_room_corridor: "lrc_", //TODO Selector for n, e, s, w
    location_room_number: "lrn",

    settings_save_btn : "rs_btn",
    room_settings_room_id : "rs_rid",
    wings : ["North", "East", "South", "West"],

    cFlrs : (e) => {
        for (let i=0;i<rapi.floors;i++) {
            const r = $$("row")
            r.classList.add("justify-content-center", "mt-5", "mb-2")
            r.innerHTML = '<div class="col-12 ros-bg-db text-white mt-2 mb-3 text-center" style=" ' +
                'border-style: solid;border-color:#000000;' +
                'border-radius: 1em" id="f'+ i+ '" onclick="rapi.selectFloor(this, true)"' +
                'onmouseenter="if(rapi.f!=this)this.style.backgroundColor = color_sceme.sr"' +
                'onmouseleave="if(rapi.f!=this)this.style.backgroundColor = color_sceme.db"' +
                '>Floor '+ i + ' </div>'
            e.appendChild(r)
        }
    },
    cWgs : (e) => {
        rapi.wings.forEach((i) => {
            const r = $$("row")
            r.classList.add("justify-content-center", "mt-5", "mb-2")
            r.innerHTML = '<div class="col-12 ros-bg-db text-white mt-2 mb-3 text-center" style=" ' +
                'border-style: solid;border-color:#000000;' +
                'border-radius: 1em" id="w'+ i+ '" onclick="rapi.selectWing(this, true)"' +
                'onmouseenter="if(rapi.w!=this)this.style.backgroundColor = color_sceme.sr"' +
                'onmouseleave="if(rapi.w!=this)this.style.backgroundColor = color_sceme.db"' +
                '>'+ i + ' </div>'
            e.appendChild(r)
        })
    } ,
    cRms : (j) =>{
        j["rooms"].forEach((e) =>{
            rapi.createRoom(e)
        })
    },
    setLevel : (i, a)=> {
        rapi.sf = Number.parseInt(i.id.substr(1,2))
        if (!a)return
        rapi.fetch(rapi.sf, rapi.sw , (j) => {
            rapi.cRms(j)
        })
    },
    setWing : (e, a) => {
        rapi.sw = e.id.substr(1,1)
        if (!a)return
        rapi.fetch(rapi.sf, rapi.sw , (j) => {
            rapi.cRms(j)
        })
    },
    init : () => {
        let l = false
        const e = $(rapi.btn)
        e.onmouseenter = () => {
            if (l)return
            e.style.backgroundColor = color_sceme.sr
        }
        e.onmouseleave = () => {
            if (l)return
            e.style.backgroundColor = color_sceme.lr
        }
        e.onclick = () => {
            if (l)return
            rapi.createNewRoom((f)=>{
                if (f)l = false;
            })
            e.style.backgroundColor = color_sceme.dr
            l = true
            setTimeout(() => {
                e.style.backgroundColor = color_sceme.lr
            },500)
        }
        $(rapi.settings_ids + rapi.room_bookable).onclick = () => {
            if (rapi.r != null){
                settings.room_bookable = !settings.room_bookable;
                if (settings.room_bookable){
                    $(rapi.settings_ids + rapi.room_bookable).style.backgroundColor = '#80F530';
                }
                else{
                    $(rapi.settings_ids + rapi.room_bookable).style.backgroundColor = color_sceme.sr;
                }
            }
        }
        $(rapi.settings_save_btn).onclick = () => {
            rapi.save()
        }
        rapi.cFlrs($(rapi.fc))
        rapi.cWgs($(rapi.wc))
        rapi.selectFloor($("f0"), false)
        $("f0").style.backgroundColor = color_sceme.dr

        rapi.selectWing($("wNorth"), false)

        rapi.fetch(rapi.sf, rapi.sw , (j) => {
            rapi.clearRooms()
            rapi.cRms(j)
        })
    },
    fetch : (f, w, cb) => {
        const r = new XMLHttpRequest();
        r.open(http.me.fetch, http.p.room_api + "?" + http.q.location_room_floor_level + "=" + f + "&" + http.q.location_room_corridor + "=" + w, true);
        r.send();
        r.onreadystatechange = () => {
            if (r.readyState == XMLHttpRequest.DONE){
                cb(JSON.parse(r.response))
            }
        }
    },
    selectFloor : (e, a)=> {
        rapi.clearRooms()
        if (rapi.f != null && rapi.f.style != null){
            rapi.f.style.backgroundColor = color_sceme.db;
        }
        rapi.f = e;
        rapi.setLevel(e, a)
        e.style.backgroundColor = color_sceme.dr;
    },
    selectWing : (e, a) => {
        rapi.clearRooms()
        if (rapi.w != null && rapi.w.style != null){
            rapi.w.style.backgroundColor = color_sceme.db;
        }
        rapi.w = e;
        rapi.setWing(e, a)
        e.style.backgroundColor = color_sceme.dr;
    },
    clearRooms : () => {
        $(rapi.rc).innerHTML = ""
        rapi.r = null;
    },
    select : (k , e) => {
        if (rapi.r != null){
            rapi.r.style.backgroundColor = color_sceme.sb;
        }
        k.style.backgroundColor = color_sceme.pr;
        rapi.r = k;

        const r = new XMLHttpRequest();
        r.open(http.me.fetch, http.p.room_api + "?" + http.q.room_id + "=" + e, true)
        r.send()
        r.onreadystatechange = () => {
            if (r.readyState == XMLHttpRequest.DONE){
                rapi.putSettings(JSON.parse(r.response))
            }
        }
    },
    putSettings : (j) => {

        settings.max_booking_duration = j[http.q.room_max_duration]
        settings.room_id = j[http.q.room_id];
        settings.room_name = j[http.q.room_name];
        settings.room_name_misc = j[http.q.room_name_misc];
        settings.arduino_attached = j[http.q.room_arduino_available];
        settings.room_bookable = j[http.q.room_bookable];
        settings.max_people = j[http.q.room_max_capacity];
        settings.room_location.name = j[http.q.room_location][http.q.location_room_name]
        settings.room_location.floor = j[http.q.room_location][http.q.location_room_floor_level];
        settings.room_location.wing = j[http.q.room_location][http.q.location_room_corridor];
        settings.room_location.nbr = j[http.q.room_location][http.q.location_room_number];
        $(rapi.settings_ids + rapi.room_max_booking_duration).value = settings.max_booking_duration;
        $(rapi.settings_ids + rapi.room_id).innerText = settings.room_id;
        $(rapi.settings_ids + rapi.room_name).value = settings.room_name;
        $(rapi.settings_ids + rapi.room_name_misc).value = settings.room_name_misc;
        $(rapi.settings_ids + rapi.room_arduino_available).value = settings.arduino_attached;//done serverside
        if (settings.room_bookable){
            $(rapi.settings_ids + rapi.room_bookable).style.backgroundColor = '#80F530';
        }
        else{
            $(rapi.settings_ids + rapi.room_bookable).style.backgroundColor = color_sceme.sr;
        }
        $(rapi.settings_ids + rapi.location_room_name).innerText = settings.room_location.name;
        rapi.putCorridor(settings.room_location.wing);
        $(rapi.settings_ids + rapi.location_room_floor_level).value = settings.room_location.floor;
        $(rapi.settings_ids + rapi.location_room_number).value = settings.room_location.nbr;
        $(rapi.settings_ids + rapi.room_max_capacity).value = settings.max_people;
    },

    save : () => {
        settings.max_booking_duration = $(rapi.settings_ids + rapi.room_max_booking_duration).value;
        console.log($(rapi.settings_ids + rapi.room_max_booking_duration));
        settings.room_name = $(rapi.settings_ids + rapi.room_name).value;
        settings.room_name_misc = $(rapi.settings_ids + rapi.room_name_misc).value;
        settings.room_location.floor = $(rapi.settings_ids + rapi.location_room_floor_level).value;
        settings.room_location.nbr = $(rapi.settings_ids + rapi.location_room_number).value;
        settings.room_location.wing = rapi.checkSettingsCorridor();
        settings.max_people = $(rapi.settings_ids + rapi.room_max_capacity).value;
        settings.room_location.nbr = $(rapi.settings_ids + rapi.location_room_number).value;
        settings.max_people =  $(rapi.settings_ids + rapi.room_max_capacity).value;
        settings.room_location.name  = rapi.buildRoomName(settings.room_location);
        $(rapi.settings_ids + rapi.location_room_name).innerText = settings.room_location.name;
        $(settings.room_id + rapi.room_name).innerText = settings.room_name;
        $(settings.room_id + rapi.room_name_misc).innerText = settings.room_name_misc;
        $(settings.room_id + rapi.location_room_name).innerText = settings.room_location.name;
        $(settings.room_id + "lrn_1").innerText = settings.room_location.name;
        $(settings.room_id + rapi.room_max_capacity).innerText = settings.max_people;
        if (settings.room_bookable){
            $(settings.room_id + rapi.room_bookable).innerText = "true"
        }else{
            $(settings.room_id + rapi.room_bookable).innerText = "false"
        }
        rapi.sendQuery();
    },
    sendQuery : () => {
        const r = new XMLHttpRequest()
        r.open(http.me.put, http.p.room_api + rapi.putQuery(), true)
        r.send()
        console.log(r)
        r.onreadystatechange = (e) => {
            if (r.readyState == XMLHttpRequest.DONE){
                const j = JSON.parse(r.response);
                console.log(j)
                if (j["success"]){
                    $(rapi.settings_save_btn).style.backgroundColor = '#80F530';
                    setTimeout(() => {
                        $(rapi.settings_save_btn).style.backgroundColor = color_sceme.lr;
                    },800);
                    console.log(j)

                    rapi.putSettings(j)
                }
            }
        }
    },
    putQuery : () => {
        return "?" +
            http.q.room_id + "=" + settings.room_id +"&"+
            http.q.room_name + "=" + settings.room_name +"&"+
            http.q.room_name_misc + "=" + settings.room_name_misc +"&"+
            http.q.room_bookable + "=" + settings.room_bookable +"&"+
            http.q.location_room_name + "=" + settings.room_location.name +"&"+
            http.q.location_room_floor_level + "=" + settings.room_location.floor +"&"+
            http.q.location_room_corridor + "=" + settings.room_location.wing +"&"+
            http.q.location_room_number + "=" + settings.room_location.nbr +"&"+
            http.q.room_max_capacity + "=" + settings.max_people + "&" +
            http.q.room_max_duration + "=" + settings.max_booking_duration
    },
    checkSettingsCorridor : () => {
        if ($(rapi.settings_ids + rapi.location_room_corridor + "n").checked){
            return 1
        }
        if ($(rapi.settings_ids + rapi.location_room_corridor + "e").checked){
            return 2
        }
        if ($(rapi.settings_ids + rapi.location_room_corridor + "s").checked){
            return 3
        }
        if ($(rapi.settings_ids + rapi.location_room_corridor + "w").checked){
            return 4
        }
        return 0
    },
    buildRoomName : (l) => {
        if (rapi.r==null)return;
        return "Z" + (l.wing==1?"N":(l.wing==2?"E":(l.wing==3?"S":(l.wing==4?"W":"_")))) + "-"+ l.floor + "-" + l.nbr
    },
    putCorridor : (e) => {
        if (e==1){
            $(rapi.settings_ids + rapi.location_room_corridor + "n").checked = true
            $(rapi.settings_ids + rapi.location_room_corridor + "e").checked = false
            $(rapi.settings_ids + rapi.location_room_corridor + "s").checked = false
            $(rapi.settings_ids + rapi.location_room_corridor + "w").checked = false
        }
        if (e==2){
            $(rapi.settings_ids + rapi.location_room_corridor + "n").checked = false
            $(rapi.settings_ids + rapi.location_room_corridor + "e").checked = true
            $(rapi.settings_ids + rapi.location_room_corridor + "s").checked = false
            $(rapi.settings_ids + rapi.location_room_corridor + "w").checked = false
        }
        if (e==3){
            $(rapi.settings_ids + rapi.location_room_corridor + "n").checked = false
            $(rapi.settings_ids + rapi.location_room_corridor + "e").checked = false
            $(rapi.settings_ids + rapi.location_room_corridor + "s").checked = true
            $(rapi.settings_ids + rapi.location_room_corridor + "w").checked = false
        }
        if (e==4){
            $(rapi.settings_ids + rapi.location_room_corridor + "n").checked = false
            $(rapi.settings_ids + rapi.location_room_corridor + "e").checked = false
            $(rapi.settings_ids + rapi.location_room_corridor + "s").checked = false
            $(rapi.settings_ids + rapi.location_room_corridor + "w").checked = true
        }
    },
    selectC : (i) => {
        if (rapi.r==null)return;
        settings.room_location.wing = i;
        settings.room_location.name  = rapi.buildRoomName(settings.room_location)
        $(rapi.settings_ids + rapi.location_room_name).innerText = settings.room_location.name;
    },
    selectRN : (i) => {
        if (rapi.r==null)return;
        settings.room_location.nbr = i;
        settings.room_location.name  = rapi.buildRoomName(settings.room_location);
        $(rapi.settings_ids + rapi.location_room_name).innerText = settings.room_location.name;
    },
    selectFN : (i) => {
        if (rapi.r==null)return;
        settings.room_location.floor = i;
        settings.room_location.name  = rapi.buildRoomName(settings.room_location);
        $(rapi.settings_ids + rapi.location_room_name).innerText = settings.room_location.name;
    },
    createRoom : (e) => {
        const r = $$("row")
        r.classList.add("justify-content-center", "mt-5", "mb-2")
        r.id = e[http.q.room_id];
        r.innerHTML = '<div class="col-12 ros-bg-sb mt-2 mb-3 text-center"' +
            'onclick="rapi.select(this, `'+ e[http.q.room_id] +'`)" ' +
            'style="border-style: solid;' +
            'border-radius: 1em"><div class="row justify-content-center"><div class="col-11 mt-2 mb-2">' +
            '<div class="row justify-content-end"><div class="col-8" id="'+e[http.q.room_id]+'lrn_1"><b><u>' + e[http.q.room_location][http.q.location_room_name] +
            '</u></b></div><button onclick="rapi.remove(`' + e[http.q.room_id]  + '`)" " class="col-3 btn ros-bg-sr text-center" ' +
            'style="width: 2em;height: 2em;font-size: 12px;padding-bottom: 1.6em;"><b>X</b></button></div>'+
            rapi.ce(e[http.q.room_id]+"rid", e, http.q.room_id, "Room ID") +
            rapi.ce(e[http.q.room_id]+"rnm",e, http.q.room_name, "Room Name") +
            rapi.ce(e[http.q.room_id]+"rmn",e, http.q.room_name_misc, "Room Misc Name") +
            rapi.ce(e[http.q.room_id]+"raa",e, http.q.room_arduino_available, "Arduino Available") +
            rapi.ce(e[http.q.room_id]+"rbb",e, http.q.room_bookable, "Room Bookable") +
            rapi.ce(e[http.q.room_id]+"lrnm",e[http.q.room_location], http.q.location_room_name, "Room Location Name") +
            rapi.ce(e[http.q.room_id]+"rmx",e, http.q.room_max_capacity, "Max Capacity") +
            '</div></div></div>'
        $(rapi.rc).appendChild(r)
    },
    createNewRoom : (f) => {
        const r = new XMLHttpRequest();
        r.open(http.me.post, http.p.room_api + "?" + http.q.location_room_floor_level +"=" + rapi.sf + "&" + http.q.location_room_corridor + "=" + rapi.sw, true)
        r.send();
        r.onreadystatechange = () => {
            if (r.readyState == XMLHttpRequest.DONE){
                //TODO generate Data on server
                const j = JSON.parse(r.response);
                if (j["success"]){
                    rapi.createRoom(j)
                    if (f!=undefined)f(true)
                }
                else{
                    if (f!=undefined)f(false)
                }
            }
        }
        /**
         * TODO
         * DO the server request, create a new room, fetch the server side room id from the database
         * implement the room id into the new container, unchangeable parameter from the clientside
         * make other variables changeable and interactable
         */
    },
    ce : (i, e , n, h) => {
        return '<div class="row mt-2"><div class="col-12 text-center " ' +
            'style="font-size: 12px;"><b>' + h + '</b></div></div><div ' +
            'class="row justify-content-center"><div class="col-12 text-center text-white ros-bg-db" ' +
            'style="border-style: solid;border-color: #000000;border-radius: 5px" id="'+ i + '">' + e[n] + '</div></div>'
    },
    remove : (i) => {
        const r = new XMLHttpRequest()
        r.open(http.me.delete, http.p.room_api + "?" + http.q.room_id + "=" + i, true)
        r.send()
        r.onreadystatechange = (e) => {
            if (r.readyState == XMLHttpRequest.DONE){
                if (JSON.parse(r.response)["success"]){
                    $(i).remove()
                }
            }
        }
    }
}

window.onload = () => {
    rapi.init()
}



