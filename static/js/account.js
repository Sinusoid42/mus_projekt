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

 This js file manages the user account page functionalities encapsulated within the
 acc const object

 This file enables changing of bookings using a simple form from the user account page, bookings can be selected
 and changed

 Upgrades can be submitted and in turn possibly be answered by a admin account priviledged user able to access the api upgrades page

 @author ben
 */





const acc = {
    bl : "booking_list",

    a : [-1,0,1,2,3,4,5],
    b : ["System-Admin", "Public", "TH-Student", "TH-Tutor", "TH-Professor", "TH-Professor-A", "TH-Room-Admin"],

    d : ["Mo", "Tue", "WE", "Thu", "Fr"],
    m : ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dez"],


    tag_booking_settings_id : "bs_id",              //Booking id, inner Text
    tag_booking_settings_room_id : "bs_rid",        //Room id, inner Text
    tag_booking_settings_topic : "bs_tpc",          //Booking Topic, value => text input
    tag_booking_settings_reoccurring : "bks_t_r",   //button, tracking by acc.settings.reoccuring
    tag_booking_settings_description : "bs_dsc",    //Description
    tag_booking_settings_count : "bs_c",
    tag_booking_settings_end_date : "bs_d",         //Booking Settings Date, returns a RFC encoded date string
    tag_booking_settings_save_button : "bs_sv",
    selected : false,

    settings: {

        reoccuring : false,
        enddate : {
            year : 0,
            month : 0,
            day : 0,
            week : 0,
        },


        booking_id : "",
        room_id : "",
        topic : "",
        description : "",
        count : 0,




    },

    upgrade : (e) => {
        console.log(e)
        const r = new XMLHttpRequest();
        r.open(http.me.post, http.p.users, true);
        r.send();
        r.onreadystatechange = () => {
            if (r.readyState == XMLHttpRequest.DONE) {
                console.log(r.response)
                const j = JSON.parse(r.response)
                if (j[http.q.success]) {
                    $("ubtn").remove()
                }
            }
        }

    },

    init : () => {
        console.log("Init the acc js")


            new Promise ((cb, er) => {
                const a = $(acc.tag_booking_settings_reoccurring)
                if (a != undefined){
                    cb(a)
                }
            }).then((e)=> {
                e.onclick = () => {
                    acc.putReoccuringData(!acc.settings.reoccuring)
                }
            })
        $(acc.tag_booking_settings_save_button).onclick = () => {
            acc.saveBK()
        }
        acc.fetch();
    },

    fetch : () => {
        $(acc.bl).innerHTML = ``
        const r = new XMLHttpRequest();
        r.open(http.me.fetch, http.p.users, true);
        r.send();
        r.onreadystatechange = () => {
            if (r.readyState == XMLHttpRequest.DONE) {
                console.log(r.response)
                const j = JSON.parse(r.response)
                if (j[http.q.success]) {
                    acc.cbks(j)
                }
            }
        }
    },
    cbks : (e) => {
        let a = ""
        e["bookings"].forEach((e) => {
            a += acc.cbk(e)
        })
        $(acc.bl).innerHTML = a;
    },

    cbk : (e) => {
        let d = acc.weekDate(e[http.q.booking_time_year], e[http.q.booking_time_week]-1)

        d.setHours(24*(e[http.q.booking_time_day]))

        if (e[http.q.booking_time_reoccuring]){
            let q = acc.weekDate(e[http.q.booking_time_end_year], e[http.q.booking_time_end_week]-1)

            q.setHours(24*(e[http.q.booking_time_end_day]))


            const l = q.toLocaleString().split("/")
            const p = l[2].split(", ")


            return      `<div class="row justify-content-center mt-1">
                        <div class="col-10 justify-content-around u_bks_le" style="border-radius: 0.2em 0.2em 0.2em 0.2em; border-color: #000000; border-style: solid;border-width: 1px;"
                            onclick="acc.selectBK('`+ e[http.q.booking_id] + `')">
                            <div class="row justify-content-around mt-2 mb-2">
                                <div class="col-5 font-weight-bold" style="border-radius: 0.2em 0.2em 0.2em 0.2em; border-color: #000000; border-style: solid;border-width: 1px;"> Booking ID </div>
                                <div class="col-5" style="border-radius: 0.2em 0.2em 0.2em 0.2em; border-color: #000000; border-style: solid;border-width: 1px;">`+ e[http.q.booking_id] + `</div>
                            </div>
                            <div class="row justify-content-around mt-2 mb-2">
                                <div class="col-5 font-weight-bold" style="border-radius: 0.2em 0.2em 0.2em 0.2em; border-color: #000000; border-style: solid;border-width: 1px;"> Room Location </div>
                                <div class="col-5" style="border-radius: 0.2em 0.2em 0.2em 0.2em; border-color: #000000; border-style: solid;border-width: 1px;">`+ e[http.q.location_room_name] + `</div>
                            </div>
                            <div class="row justify-content-around mt-2 mb-2">
                                <div class="col-5 font-weight-bold" style="border-radius: 0.2em 0.2em 0.2em 0.2em; border-color: #000000; border-style: solid;border-width: 1px;"> Booking Topic </div>
                                <div class="col-5" style="border-radius: 0.2em 0.2em 0.2em 0.2em; border-color: #000000; border-style: solid;border-width: 1px;">`+ e[http.q.booking_topic] + `</div>
                            </div>
                            <div class="row justify-content-around mt-2 mb-2">
                                <div class="col-5 font-weight-bold" style="border-radius: 0.2em 0.2em 0.2em 0.2em; border-color: #000000; border-style: solid;border-width: 1px;"> Date </div>
                                <div class="col-5" style="border-radius: 0.2em 0.2em 0.2em 0.2em; border-color: #000000; border-style: solid;border-width: 1px;">`+acc.d[d.getDay()-1] + ` : ` + (d.getUTCDate()+1) + ` - ` + acc.m[(d.getMonth())] + ` - `+ d.getFullYear() + `</div>
                            </div>
                            <div class="row justify-content-around mt-2 mb-2">
                                <div class="col-5 font-weight-bold" style="border-radius: 0.2em 0.2em 0.2em 0.2em; border-color: #000000; border-style: solid;border-width: 1px;"> End Date </div>
                                <div class="col-5" style="border-radius: 0.2em 0.2em 0.2em 0.2em; border-color: #000000; border-style: solid;border-width: 1px;">`+acc.d[q.getDay()-1] + ` : ` + (parseInt(l[0])) + ` - ` + acc.m[(parseInt(l[1])-1)] + ` - `+ p[0] + `</div>
                            </div>
                        </div>
                    </div>`
        }
        return      `<div class="row justify-content-center mt-1">
                        <div class="col-10 justify-content-around u_bks_le" style="border-radius: 0.2em 0.2em 0.2em 0.2em; border-color: #000000; border-style: solid;border-width: 1px;"
                        onclick="acc.selectBK('`+ e[http.q.booking_id] + `')">
                            <div class="row justify-content-around mt-2 mb-2">
                                <div class="col-5 font-weight-bold" style="border-radius: 0.2em 0.2em 0.2em 0.2em; border-color: #000000; border-style: solid;border-width: 1px;"> Booking ID </div>
                                <div class="col-5" style="border-radius: 0.2em 0.2em 0.2em 0.2em; border-color: #000000; border-style: solid;border-width: 1px;">`+ e[http.q.booking_id] + `</div>
                            </div>
                            <div class="row justify-content-around mt-2 mb-2">
                                <div class="col-5 font-weight-bold" style="border-radius: 0.2em 0.2em 0.2em 0.2em; border-color: #000000; border-style: solid;border-width: 1px;"> Room Location </div>
                                <div class="col-5" style="border-radius: 0.2em 0.2em 0.2em 0.2em; border-color: #000000; border-style: solid;border-width: 1px;">`+ e[http.q.location_room_name] + `</div>
                            </div>
                            <div class="row justify-content-around mt-2 mb-2">
                                <div class="col-5 font-weight-bold" style="border-radius: 0.2em 0.2em 0.2em 0.2em; border-color: #000000; border-style: solid;border-width: 1px;"> Booking Topic </div>
                                <div class="col-5" style="border-radius: 0.2em 0.2em 0.2em 0.2em; border-color: #000000; border-style: solid;border-width: 1px;">`+ e[http.q.booking_topic] + `</div>
                            </div>
                            <div class="row justify-content-around mt-2 mb-2">
                                <div class="col-5 font-weight-bold" style="border-radius: 0.2em 0.2em 0.2em 0.2em; border-color: #000000; border-style: solid;border-width: 1px;"> Date </div>
                                <div class="col-5" style="border-radius: 0.2em 0.2em 0.2em 0.2em; border-color: #000000; border-style: solid;border-width: 1px;">`+acc.d[d.getDay()-1] + ` : ` + (d.getUTCDate()+1) + ` - ` + acc.m[(d.getMonth())] + ` - `+ d.getFullYear() + `</div>
                            </div>
                        </div>
                    </div>`
    },

    selectBK : (e) => {

        const r = new XMLHttpRequest()
        r.open(http.me.select, http.p.users + "?" + http.q.booking_id + "=" + e, true)
        r.send()
        r.onreadystatechange = () => {
            if (r.readyState == XMLHttpRequest.DONE){
                const j = JSON.parse(r.response)
                console.log(j)
                if (j[http.q.success]){

                    acc.settings.booking_id = j[http.q.booking_id]
                    acc.settings.room_id = j[http.q.booking_room_id]

                    $(acc.tag_booking_settings_id).innerText = j[http.q.booking_id]
                    $(acc.tag_booking_settings_room_id).innerText = j[http.q.booking_room_id]
                    $(acc.tag_booking_settings_count).value = j[http.q.booking_occupancy]
                    $(acc.tag_booking_settings_description).value = j[http.q.booking_description]
                    $(acc.tag_booking_settings_topic).value = j[http.q.booking_topic]
                    acc.putReoccuringData(j[http.q.booking_time_reoccuring])
                    if (j[http.q.booking_time_reoccuring]){
                        acc.putEndDate(j)
                    } else{
                        acc.clearEndDate()
                    }

                    new Promise((f) => {
                        const b = $(acc.tag_booking_settings_reoccurring)
                        if (b != undefined){
                            f(b)
                        }
                    }).then((b)=> {
                        b.removeAttribute("disabled", true)
                    })
                    $(acc.tag_booking_settings_save_button).removeAttribute("disabled", true)

                }
            }
        }
    },

    saveBK : () => {
      //acc.settings.reoccuring contains already proper bool value
        acc.settings.count = $(acc.tag_booking_settings_count).value
        acc.settings.description = $(acc.tag_booking_settings_description).value
        acc.settings.topic = $(acc.tag_booking_settings_topic).value
        let q = "?" +
            http.q.booking_id + "=" + acc.settings.booking_id + "&" +
            http.q.room_id + "=" + acc.settings.room_id + "&" +
            http.q.booking_description + "=" + acc.settings.description + "&" +
            http.q.booking_occupancy + "=" + acc.settings.count + "&" +
            http.q.booking_topic + "=" + acc.settings.topic
        if (acc.settings.reoccuring) {
                const e = $(acc.tag_booking_settings_end_date)
                if (e != undefined) {
                const date = new Date(parseInt(e.value.split("-")[0]), parseInt(e.value.split("-")[1]) - 1, parseInt(e.value.split("-")[2])-1, 0, 0, 0, 0)
                const oneJan = new Date(date.getFullYear(), 0, 1);
                const ew = Math.ceil((date.valueOf() - oneJan.valueOf()) / (24 * 60 * 60 * 1000) / 7);
                q +=  "&" + http.q.booking_time_reoccuring + "=" + acc.settings.reoccuring + "&" +
                    http.q.booking_time_end_year + "=" + date.getFullYear() + "&"+
                    http.q.booking_time_end_month + "=" + (date.getMonth()+1) + "&"+
                    http.q.booking_time_end_day + "=" + (date.getDay()) + "&"+
                    http.q.booking_time_end_week + "=" + ew
                }else{
               q +=  "&" + http.q.booking_time_reoccuring + "=" + false + "&" +
                    http.q.booking_time_end_year + "=" + 0 + "&"+
                    http.q.booking_time_end_month + "=" + 0 + "&"+
                    http.q.booking_time_end_day + "=" + 0 + "&"+
                    http.q.booking_time_end_week + "=" + 0
                }
        }
        const r = new XMLHttpRequest()
        r.open(http.me.put, http.p.users + q , true)
        r.send()
        r.onreadystatechange = () => {
            if (r.readyState == XMLHttpRequest.DONE){
                const j = JSON.parse(r.response)
                if (j[http.q.success]){
                    acc.fetch()
                }
            }
        }
    },

    putReoccuringData : (e) => {
        console.log(e)
        acc.settings.reoccuring = e

        new Promise((cb)=> {
            const b = $(acc.tag_booking_settings_reoccurring)
            if (b != undefined){
                cb(b)
            }
        }).then((b)=> {
            if (e) {
                b.classList.add("u_brs_r_1")
                b.classList.remove("u_brs_r_0", "u_bks_r")
                b.removeAttribute("disabled", true)
                $(acc.tag_booking_settings_save_button).removeAttribute("disabled", true)
                console.log($(acc.tag_booking_settings_save_button))
                console.log(b, e)
            }else{
                console.log(b)
                b.removeAttribute("disabled", true)
                console.log(b)
                $(acc.tag_booking_settings_save_button).removeAttribute("disabled", true)
                b.classList.add("u_brs_r_0")
                b.classList.remove("u_brs_r_1", "u_bks_r")
            }
        })


    },
    putEndDate : (j) => {
        console.log(j)
        let d = acc.weekDate(j[http.q.booking_time_end_year], j[http.q.booking_time_end_week]-1)

        d.setHours(24*(j[http.q.booking_time_end_day]))

        const a = d.toLocaleString().split("/")
        const b = a[2].split(", ")
        new Promise((f) => {
            const e = $(acc.tag_booking_settings_end_date)
            if (e != undefined){
                f(e)
            }
        }).then((e)=> {
            e.value = b[0] + "-"+ a[1] + "-" + a[0]
        })

    },

    clearEndDate : () => {
        new Promise((c) => {
            const e = $(acc.tag_booking_settings_end_date)
            if (e!= undefined){
                c(e)
            }
        }).then((e)=> {
            e.value = "2021-01-01"
        })

    },

    /**
     * Returns a new Date Object according to
     * the year and week given
     * the Date represents the first monday
     * of that week
     * @param y
     * @param w
     * @returns {Date}
     */
    weekDate : (y, w) => {
        let d = new Date(y,0,1,0,0,0,0);
        while(true){
            if (d.getDay() == 1){
                //monday
                break;
            }
            d.setHours(d.getHours()+24,0,0,0)
            if (d.getDay() == 1){
                //monday
                break;
            }
        }
        d.setHours(d.getHours()+w*7*24,0,0,0)
        return d;
    },




}