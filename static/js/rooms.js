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

 Room Logic Implementations

 This JS file focusses on serving all
 functionalities when working with the pubilc rooms.tmpl
 @author benwinter (017.07.2021)
    -logic and local storage
    -data transfer and definitions
    -pipeline integration
 @author davidmartinkarg (028.07.2021)
    -init
    -calender setup
    -graphical design
    -layout
 */


/**
 * The calender namespace,
 * handles the events given by the calender frontend
 *
 * applies async data to the layout, so the frontend
 * is pretty much as responsive, without css on focus,
 * but in focus of data transfer, communication and
 * design
 * @type {{THURSDAY: number, select: cal.select, year: number, getCalenderWeek: (function(): *|string), getCurrentYear: (function(): number), WEDNESDAY: number, deselectColumn: cal.deselectColumn, ccol: cal.ccol, c_b_btn: string, repeating: boolean, clearSelection: cal.clearSelection, SATURDAY: number, checkSelectedColumn: cal.checkSelectedColumn, weekDate: (function(*=, *): Date), randomize: cal.randomize, selected: boolean, requestNewBooking: cal.requestNewBooking, init: cal.init, SUNDAY: number, rpbtn: string, b_c_m: number, TUESDAY: number, b_c_w: number, selected_column_min: number, addSelected: cal.addSelected, selected_column: number, checkSelectedData: (function(*): *), checkRemoval: (function(*, *): boolean), selected_column_max: number, t1: string[], t2: string[], room_id: string, s_b_id: string, noc: number, week: number, removeSelected: cal.removeSelected, putSettingsData: cal.putSettingsData, FRIDAY: number, update: cal.update, cell: cal.cell, getCalenderYear: (function(): *|string), nor: number, b_c_y: number, getCalenderValue: (function(*, *): string), cc: cal.cc, c_ct: string, ccol_t: cal.ccol_t, selected_ids: [], cm: string, r_d_id: string, MONDAY: number, u: cal.u, fetch: cal.fetch, getCurrentWeek: (function(): number), selectColumn: cal.selectColumn, s_b_rid: string, createCreateQuery: cal.createCreateQuery, putCalenderData: cal.putCalenderData}}
 */
const cal={
    /**
     * namespace variables and ids
     */
    nor: 14,        //number of rows
    noc: 5,         //number of cols
    b_c_y:0,        //booking current year
    b_c_m:0,        //booking current month
    b_c_w:0,        //booking current week
    c_ct: "c_ct",   //calender container

    m_b_d : 0,      //max booking duraction allowed by the server, when longer periods are given, the server starts at the beginning
    r_c : 0,        //room_capacity
    r_o: 0,         //room_occupancy
    /**
     * Days of the week as integer
     */
    MONDAY : 1,
    TUESDAY : 2,
    WEDNESDAY : 3,
    THURSDAY : 4,
    FRIDAY : 5,
    SATURDAY : 6,   //deprecated
    SUNDAY : 7,     //deprecated



    lb : true,      //listing bookings or settings


    /**
     * timers for the current calender
     */
    year : 0,
    week : 0,
    month : 0,
    day : -1,
    /**
     * Calender Menu
     */
    cm : "c_cm",              //CALENDER SELECTOR

    /**
     * Settings
     */
    c_b_btn : "c_b_btn",    //create booking button

    /**
     * settings booking id
     * settins room id
     */
    rpbtn : "rpbtn",       //repeat button

    s_b_rptd : "s_b_rptd", //repeat day selection
    s_b_rptd_d : "s_b_rptd_d", //repeat day selection textfield
    s_b_ed : "s_b_ed",      //settings booking end date

    s_b_id : "s_b_id",      //settings booking id => gets the id for the booking, when the post creates a new booking for the room
    s_b_rid : "s_b_rid",    //settings booking room id => The room id for the current selection, when the booking is done
    s_b_t : "s_b_t",        //settings booking topic
    s_b_c : "s_b_c",        //settings booking count
    s_b_d : "s_b_d",        //settings booking description

    r_m_c : "r_m_c",        //room max capacity
    r_c_o : "r_c_o",        //room current occupancy


    room_id : "",

    bkf : "booking_form",
    auth : false,

    /**
     * helper vars
     */
    selected_ids : [],
    selected : false,
    selected_column_min : 0,
    selected_column_max : 0,
    selected_column : 0,
    selected_occupancy : 0,
    repeating : false,
    capacity : 0,

    d : ["MO", "TUE", "WE", "THU", "FR"],
    t1:["8:00","8:50","9:45","10:35","11:30","12:20","13:15","14:05","15:00","15:50","16:45","17:35","18:30","19:20" ],
    t2:["8:45","9:35","10:30","11:20","12:15","13:05","14:00","14:50","15:45","16:35","17:30","18:20","19:15","20:05" ],

    t3:[8   ,8  ,9  ,10 ,11 ,12 ,13 ,14 ,15 ,15 ,16 ,17 ,18 ,19],
    t4:[0   ,50 ,45 ,35 ,30 ,20 ,15 ,5  ,0  ,50 ,45 ,35 ,30 ,20],
    t5:[8   ,9  ,10 ,11 ,12 ,13 ,14 ,14 ,15 ,16 ,17 ,18 ,19 ,20],
    t6:[45  ,35 ,30 ,20 ,15 ,5  ,0  ,50 ,45 ,35 ,30 ,20 ,15 ,5],

    data : {},

    /**
     * Initializes the calender frontend
     * fetches data into the localstorage to
     * generate a new calenderweek from
     * the revicived data and performs
     * initialization checks
     */
    init:()=>{
        cal.room_id = rid;
        cal.cd(cal.nor, cal.noc)


        cal.cc(cal.nor, cal.noc);
        //cal.randomize(data, cal.nor, cal.noc);
        cal.update(cal.data, cal.nor, cal.noc);

        cal.week = cal.getCurrentWeek();
        cal.year = cal.getCurrentYear();
        cal.month = cal.getCurrentMonth();

        $(cal.cm).value = cal.getCalenderValue(cal.year, cal.week)
        cal.paintCurrentHour();

        /**
         * Callback setup
         */

        $(cal.cm).onchange = () => {
            cal.week = cal.getCalenderWeek();
            cal.year = cal.getCalenderYear();
            cal.month = cal.getCalenderMonth();
            cal.clearSelection();
            cal.fetch(cal.room_id, cal.year, cal.week)

        }

        cal.fetch(cal.room_id, cal.year, cal.week)
        /**cal.putCalenderData(data, data, cal.nor, cal.noc)
        cal.randomize(data, cal.nor, cal.noc)
        cal.u()*/
    },


    /**
     * creates the clientsite data structure
     */
    cd : (r, c) => {
        for (let i=0;i<c; i++) {
            cal.data["m" + i] = {}
            for (let j = 0; j < r; j++) {
                cal.data["m" + i]["m" + i + "_" + j] = {
                    booking_id: "",
                    occupancy: 0,
                    occupied: false,
                    bookable: false,
                    selected: false,
                }
            }
        }
    },

    /**
     * Returns the current Week of @LocalGeoLocation for
     * the client, that way in terms of usage of different
     * timezones on weekends, the server will respond appropriately
     * according to the timezone week, the client is lcoated
     * @returns {number}
     */
    getCurrentWeek : () => {
        const date = new Date()
        const oneJan = new Date(date.getFullYear(),0,1);
        const numberOfDays = Math.ceil((date.valueOf() - oneJan.valueOf()) / (24 * 60 * 60 * 1000) / 7);
        return numberOfDays
    },

    /**
     * Returns the current Year of @LocalGeoLocation
     * @returns {number}
     */
    getCurrentYear : () => {
        return new Date().getFullYear();
    },

    getCurrentMonth : () => {
        return new Date().getMonth()+1;
    },


    /**
     * Returns a value useful for setting the DOMElement.value of the
     * DOM calender input element according to a set value
     * @param i
     * @param j
     * @returns {string}
     */
    getCalenderValue : (i, j) => {
        return (i + "-W" + j)
    },

    cbkf : (auth, admin_level) => {
      var p = `<div class="row mt-2"><div class=" col-12 text-center font-weight-bold" style="font-size: 18px;font-style: italic;">Create Booking
                                    </div></div><div class="row pt-2"> <div class="col-4  pl-0 pr-2"> <div class="col-12 input-group-text pt-0 pb-0 mt-0 mb-0">Booking ID</div>
                                    </div><div class="col-7  pl-0 pr-0"><div class="input-group-text  pt-0 pb-0 mt-0 mb-0" id="s_b_id">-</div>
                                    </div></div><div class="row pt-2"><div class="col-4  pl-0  pr-2"><div class="input-group-text  pt-0 pb-0 mt-0 mb-0">Room ID
                                    </div></div><div class="col-7  pl-0 pr-0"><div class="input-group-text  pt-0 pb-0 mt-0 mb-0" id="s_b_rid">`+ cal.room_id + `
                                    </div></div></div><div class="row pt-2"><div class="col-4  pl-0  pr-2"><div class="input-group-text  pt-0 pb-0 mt-0 mb-0">Topic
                                    </div></div><div class="col-7  pl-0 pr-0"><input class="input-group-text  pt-0 pb-0 mt-0 mb-0" id="s_b_t">
                                    </div></div><div class="row pt-2"><div class="col-4  pl-0  pr-2"><div class="input-group-text  pt-0 pb-0 mt-0 mb-0">Count</div>
                                    </div><div class="col-7 "><div class="row justify-content-center"><input type="number" min="0" value="0" class="col-4 input-group-text  pt-0 pb-0 mt-0 mb-0" id="s_b_c" placeholder="0">
                                    </div></div></div><div class="row pt-2"><div class="col-4  pl-0  pr-2"><div class="input-group-text  pt-0 pb-0 mt-0 mb-0">Description
                                    </div></div><textarea  class="col-7 input-group-text" rows="3" id="s_b_d"  placeholder="description"></textarea></div>`
        const d = new Date();
        const Y = d.getFullYear();
        let M = d.getMonth()+1;
        M = M<10?"0"+M:M;
        let D = d.getDay()+1;
        D = D<10?"0"+D:D;
        const dd = Y+"-"+M+"-"+D;
        if (admin_level == 2 ||
            admin_level == 3 ||
            admin_level == 4 ||
            admin_level == 5 ||
            admin_level == -1 && auth) p += `<div class="row pt-2"><div class="col-4  pl-0  pr-2"><div class="input-group-text  pt-0 pb-0 mt-0 mb-0">Repeated
                                     </div></div><div class="col-7 "><div class="row justify-content-center"><button type="date" class="btn ros-bg-sr pt-0 pb-0 mt-0 mb-0 rpbtn" id="rpbtn">X</button>
                                     </div></div></div><div class="row pt-2"><div class="col-4  pl-0  pr -2"><div class="input-group-text  pt-0 pb-0 mt-0 mb-0">WeekDay
                                     </div></div><div class="col-7 "><div class="row justify-content-center">
                                     <div class="col-4 input-group-text  pt-0 pb-0 mt-0 mb-0" id="s_b_rptd_d">-</div></div></div></div><div class="row pt-2"><div class="col-4  pl-0  pr-2">
                                     <div class="input-group-text  pt-0 pb-0 mt-0 mb-0">EndDate</div></div><input type="date" value="`+ dd+ `" class="input-group-text pt-0 pb-0 mt-0 mb-0" id="s_b_ed"></div>`
        if (auth)              p += `<div class="row justify-content-center mt-3"><button class="btn ros-bg-pb col-8 text-center cbtn" id="c_b_btn">Create Booking</button></div>`
        $(cal.bkf).innerHTML = p


        const l = $(cal.c_b_btn)
        if (l != undefined){
            l.setAttribute("disabled", true)
            l.onclick = () => {
                if (cal.selected){
                    cal.requestNewBooking(cal.room_id);
                }
            }
        }

        const j = $(cal.rpbtn)
        if (j!=undefined){
            j.setAttribute("disabled", "true")
            j.onclick = () => {
                if (cal.selected){
                    if (cal.repeating){
                        j.style.backgroundColor = ""
                        j.classList.add("rpbtn", "ros-bg-sr")
                        cal.repeating = false
                    }else{
                        j.style.backgroundColor = '#80F530';
                        j.classList.remove("rpbtn", "ros-bg-sr")
                        cal.repeating = true;
                    }
                }
            }
        }


        $(cal.s_b_t).onchange = () => {
            cal.authCreation()
        }
        $(cal.s_b_d).onchange = () => {
            cal.authCreation()
        }
        $(cal.s_b_c).onchange = () => {
            cal.authCreation()
        }




    },

    authCreation : () => {
      if (cal.selected){
          if ($(cal.s_b_t).value.length > 0 &&
              $(cal.s_b_d).value.length > 0 &&
              $(cal.s_b_c).value > 0) {

              const i = $(cal.c_b_btn);
              if (i != undefined) i.removeAttribute("disabled", true)
              const q = $(cal.rpbtn)
              if (q != undefined) q.removeAttribute("disabled", true)
              return true;
          }
      }

        const i = $(cal.c_b_btn);
        if (i != undefined) i.setAttribute("disabled", true)
        const q = $(cal.rpbtn)
        if (q != undefined) q.setAttribute("disabled", true)
      return false;
    },


     cbkl : (e) => {
        var p = ``
        e.bookings.forEach((k) => {
            p += `<div class="row mt-2 justify-content-center"><div class="col-10" style="background-color: #FFFFFF;border-style: solid;border-color: #000000;border-radius: 0.3em 0.3em 0.3em 0.3em;border-width: 1px;">`+

                                        `<div class="row mt-1"><div class="col-4 text-center font-weight-bold">User</div><div class="col-6">` + k[http.q.user_name]  +`</div>`+
                                        (k["o"]? `<div class="col-1 font-weight-bold ros-bg-dr mt-02 mb-0 pt-0 pb-0 ml-1 mr-1 pl-0 pr-0 text-center" style="border-style: solid;border-color: #000000;border-radius: 
                                0.3em 0.3em 0.3em 0.3em;border-width: 1px;" onclick="cal.deleteBooking('`+ k[http.q.booking_id]  +`')">X</div>`:``) +
                                        `</div>`+
                                        `<div class="row"><div class="col-4 text-center font-weight-bold">BookingID</div><div class="col-6">` + k[http.q.booking_id]  +`</div></div>
                                        <div class="row"><div class="col-4 text-center font-weight-bold">Topic</div><div class="col-6">` + k[http.q.booking_topic]  +`</div></div>
                                        <div class="row"><div class="col-4 text-center font-weight-bold">Start</div><div class="col-6">` + cal.t1[k[http.q.booking_time_start]]  +`</div></div>
                                        <div class="row"><div class="col-4 text-center font-weight-bold">End</div><div class="col-6">` + cal.t2[k[http.q.booking_time_end]]  +`</div></div>
                                        </div></div>`
        })

        $(cal.bkf).innerHTML = p

    },

    deleteBooking : (e) => {
        const r = new XMLHttpRequest()
        const a = http.p.rooms+"/"+cal.room_id+"?" + http.q.booking_id + "=" + e

        r.open(http.me.delete,a, true )
        r.send()
        r.onreadystatechange = () => {
            if (r.readyState == XMLHttpRequest.DONE){
                const j = JSON.parse(r.response)
                if (j[http.q.success]){
                    //deletion success


                    cal.clearSelection()
                    cal.removeSettings();
                    cal.lb = true;

                    cal.fetch(cal.room_id, cal.year, cal.week);//reload
                }
            }
        }
    },

    paintCurrentHour : () => {
        const d = new Date()
        let a = d.getDay()
        let b = d.getHours()
        const c = d.getMinutes()
        if (a>5){
            $("m"+(4)+"_"+13).style.backgroundColor = color_sceme.dr;//first
            $("m"+(4)+"_"+13).innerText = b + ":"+c
            $("m"+(4)+"_"+13).style.color = '#000000';   //pause
            return
        }
        if (b<cal.t3[0]){
            $("m"+(a-1)+"_"+0).style.backgroundColor = color_sceme.dr;//first
            $("m"+(a-1)+"_"+0).innerText = b + ":"+c
            $("m"+(a-1)+"_"+0).style.color = '#000000';   //pause
            return
        }
        if (b>=cal.t5[cal.t5.length]){
            const q = $("m"+(a-1)+"_"+(cal.nor-1))
            if (q!=null){
                q.style.backgroundColor = color_sceme.dr;//last
                q.innerText = b + ":"+c
                q.style.color = '#000000';   //pause
            }
            return
        }
        for (let i = 0;i<cal.nor;i++){

            if ((cal.t3[i]==b && c>=cal.t4[i] ) || (cal.t5[i] == b && c<=cal.t6[i]) ){

                const q = $("m"+(a-1)+"_"+i)
                if (q!= null){
                    q.style.backgroundColor = color_sceme.dr;  //current
                    q.style.color = '#000000';   //pause
                    q.innerText = b + ":"+ (c<10?"0"+c:""+c)
                }
                return
            }
            if (i>0 && b == cal.t5[i-1] && c>=cal.t6[i-1]){
                const q = $("m"+(a-1)+"_"+i)
                if (q!=null){
                    q.style.backgroundColor = '#FFFF00';   //pause
                    q.style.color = '#000000';   //pause
                    q.innerText = b + ":"+c
                }
                return
            }
        }
    },

    /**
     * Creates a mew calender layout
     * @param r number of rows
     * @param c number of cols
     */
    cc: (r,c)=>{
        cal.ccol_t(cal.t1, cal.t2);
        for(let i=0; i<c;i++) {
            cal.ccol(i, r, cal.data);
        }
        cal.paintCurrentHour();
    },


    /**
     * create time column
     * @param e0    the start time string array
     * @param e1    the end   time string array
     */
    ccol_t:(e0, e1)=>{
        const f= $$("div");
        f.classList.add("col-2");
        $(cal.c_ct).appendChild(f);
        e0.forEach((s,i)=>{
            const row=$$("div");
            row.classList.add("row", "justify-content-start", "text-center", "mt-1", "m-0", "p-0");
            //row.style.backgroundColor=color_sceme.dr
            f.appendChild(row);
            row.innerHTML = `<div class="col-12 m-0 pl-3 pr-3 p-0 text-center"><p class=" m-0 p-0 cal_e">` + s + `-` + e1[i] + `</p></div>`
        })
    },

    fetch : (room_id, year, week) => {
        const r = new XMLHttpRequest();
        r.open(http.me.fetch, http.p.rooms + "/"+room_id + "?" + http.q.calender_year + "=" + year + "&" + http.q.calender_week + "=" + week)

        r.send()
        r.onreadystatechange = (e) =>{
            if(r.readyState == XMLHttpRequest.DONE){
                const j = JSON.parse(r.response);


                if(j[http.q.success]){
                    //cal.cbkf(j[http.q.authenticated], j[http.q.access]);
                    cal.auth = j[http.q.authenticated]
                    cal.access = j[http.q.access]
                    cal.capacity = j[http.q.room_max_capacity];
                    cal.m_b_d = j[http.q.room_max_duration];
                    cal.room_id = j[http.q.booking_room_id];
                    cal.paintCurrentHour()
                    cal.putCalenderData(cal.data, j, cal.nor, cal.noc)
                    cal.putSettingsData()
                    cal.u()
                    cal.paintCurrentHour();
                }
            }
        }
    },

    createNewBooking : () => {
      cal.lb = false;
      cal.removeSettings();
      cal.rebuildSettings();
    },


    rebuildSettings : () => {
        cal.cbkf(cal.auth, cal.access);
        cal.putSettingsData()
        cal.u()
    },

    removeSettings : () => {
        $(cal.bkf).innerHTML ='';
        $(cal.r_c_o).innerText = "Occupancy: -";
    },

    /**
     * Shortform call function to cal.update
     * Incorporates already existing predefined
     * parameters for the function update
     */
    u : () => {
      cal.update(cal.data, cal.nor, cal.noc)
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
        }
        d.setHours(d.getHours()+w*7*24,0,0,0)
        return d;
    },


    putSettingsData : () => {
        const a = $(cal.r_m_c);
        const b = $(cal.s_b_rid);
        const c = $(cal.r_m_c)
        c.innerText = "Capacity: " + cal.capacity
        if (a!= undefined && b != undefined){
            a.innerText = "Capacity: " + cal.capacity;
            b.innerText = cal.room_id;
        }
    },


    /**
     * Testfunction, randomizes the testset data for the calender
     * @param d     the data
     * @param r     the rows
     * @param c     the cols
     */
    randomize : (d, r, c) => {
        for (let m=0;m<c; m++){
            for (let n=0;n<r;n++){
                d["m"+m]["m"+m+"_"+n].bookable = Math.random()<0.8;
            }
        }
    },

    /**
     * Puts the data fetched from the server in a certain pattern
     * into the data from the client
     * @param k
     * @param e
     * @param i
     * @param j
     */
    putCalenderData : (k, e, i, j) => {
        for (let m=0;m<j; m++){
            for (let n=0;n<i; n++){
                const l = e["m"+m]["m"+m+"_"+n];
                if (l != null){
                    const p = k["m"+m]["m"+m+"_"+n];

                    p.bookable = l.b;
                    p.booking_id = l.bid;
                    p.occupied = l.o;
                    p.occupancy = l.oc;
                }
            }
        }
    },

    /**
     * create column
     * @param i the column
     * @param j the row
     * @param data
     */
    ccol: (i,j, data)=>{
        const f= $$("div");
        f.classList.add("col-2");
        $(cal.c_ct).appendChild(f);
        for(let k=0; k<j;k++){
            const row=$$("div");
            row.classList.add("row", "justify-content-left", "text-center", "mt-1", "m-0", "p-0");
            //row.style.backgroundColor=color_sceme.dr
            f.appendChild(row);
            s = "-"
            row.innerHTML = `<div class="col-6 m-0 p-0 text-center"><button class=" m-0 p-0 cal_e_b" onclick="cal.select(` + i +`,`+ k +`)" id="` +"m"+i+"_"+k+ `">` + s + `</button></div>`
        }
    },

    /**
     * Update a single cell in css ruling
     * @param d
     * @param o
     */
    cell : (d, o) => {
            o.style.borderColor = "#000000";
        if (d.bookable && !d.occupied) {
            o.classList.remove("cal_e_b", "cal_e_b_o0", "cal_e_b_o1", "cal_e_b_o2", "cal_e_b_o3", "cal_e_b_o4")
            o.classList.add("cal_e_b")
        }
        if (d.bookable && d.occupied) {
            o.classList.remove("cal_e_b", "cal_e_b_o0", "cal_e_b_o1", "cal_e_b_o2", "cal_e_b_o3", "cal_e_b_o4")
            o.classList.add("cal_e_b_o1")
        }
        if (!d.bookable) {
            o.classList.remove("cal_e_b", "cal_e_b_o0", "cal_e_b_o1", "cal_e_b_o2", "cal_e_b_o3", "cal_e_b_o4")
            o.classList.add("cal_e_b_o2")
        }
        if (d.selected){
            o.classList.remove("cal_e_b", "cal_e_b_o0", "cal_e_b_o1", "cal_e_b_o2", "cal_e_b_o3", "cal_e_b_o4")
            o.classList.add("cal_e_b_o0")
        }
    },

    /**
     * Updates all cells graphically depending on the
     * status within the data
     * @param e
     * @param i
     * @param j
     */
    update : (e, i, j) => {
      for (let m=0;m<j;m++) {
          for (let n = 0; n < i; n++) {
              const d = e["m" + m]["m" + m + "_" + n]
              const o = $("m" + m + "_" + n)
              cal.cell(d, o)
          }
      }
    },

    /**
     * selects a column to be selected,
     * main chain event handling function in the pipeline
     *
     * @param c The Column
     * @param e The List element
     */
    selectColumn : (c, e) => {
        if (cal.selected){
            //TODO Checks
        }
        cal.selected = true;
        cal.selected_column_min = e;
        cal.selected_column_max = e;
        cal.selected_column = c;
        const i = $(cal.s_b_rptd_d)
        if (i!=undefined){
            i.innerText = cal.d[cal.selected_column]
        }

        cal.day = c+1;
        for (let i = 0;i<cal.noc;i++){
            for (let j=0;j<cal.nor;j++){
                if (i!=c){
                    if (cal.data["m"+i]["m"+i+"_"+j].bookable)$("m"+i +"_"+j).style.backgroundColor = color_sceme.lr
                }
            }
        }
    },

    /**
     * Appclies color rules to the cells, when interactin with the user
     * @param c
     */
    deselectColumn : (c) => {
        cal.day = -1;
        cal.r_o = 0;
        const i = $(cal.s_b_rptd_d)
        if (i!=undefined){
            i.innerText = "-"
        }
        for (let i = 0;i<cal.noc;i++){
            for (let j=0;j<cal.nor;j++){
                if (i!=c){
                    $("m"+i +"_"+j).style.backgroundColor = "";
                }
            }
        }
    },

    /**
     * Selects a current clicked on cell
     * from the calender
     *
     * initiates the main routine to check the cell
     * to css switching and clientside datamanegement
     * @param j
     * @param k
     */
    select : (j, k) => {
        if (cal.lb){
            const r = new XMLHttpRequest()
            const q = http.p.rooms+"/"+cal.room_id+"?"+http.q.calender_year+"="+cal.getCalenderYear()+"&"+http.q.calender_week+"="+cal.getCalenderWeek()+"&"+http.q.calender_day+"="+j+"&"+http.q.calender_hour+"="+k
            r.open(http.me.select, q, true)

            r.send()
            r.onreadystatechange = () => {
                if (r.readyState == XMLHttpRequest.DONE){
                    const j = JSON.parse(r.response)
                    if (j[http.q.success]){
                        cal.putNewBookingData(j);

                    }else{
                        cal.removeSettings()
                    }
                }
            }
        }
        else {
            const o = $("m" + j + "_" + k)
            var d = cal.data["m" + j]["m" + j + "_" + k]

            if (cal.selected) {
                //TODO Check
                if (j != cal.selected_column) {
                    return
                }
                if (d.selected) {
                    if (cal.checkRemoval(j, k)) {
                        cal.removeSelected(j, k)
                        cal.cell(d, o)
                    }
                } else {
                    if (cal.checkSelectedColumn(j, k)) {
                        cal.addSelected(d, j, k)
                        cal.cell(d, o)
                        cal.authCreation()
                    }
                }
                return
            }
            if (cal.checkSelectedData(d)) {

                cal.selectColumn(j, k)
                cal.addSelected(d, j, k)
                cal.cell(d, o);
                cal.authCreation()
            }
        }
    },


    /**
     * serverside post handling => client post method is responeded with new
     * unique id, that is used for the booking
     * @param room_id
     * @param j
     * @param k
     */
    requestNewBooking : (room_id) => {
      const r = new XMLHttpRequest();

      const q = cal.createCreateQuery();

        r.open(http.me.post, http.p.rooms + "/" + room_id + q, true);
        r.send()
        r.onreadystatechange = (e) => {
            if (r.readyState == XMLHttpRequest.DONE){
                const j = JSON.parse(r.response)
                if (j[http.q.success]){

                    cal.lb = true;
                    cal.clearSelection()
                    cal.removeSettings();
                    cal.putNewBookingData(j);

                    cal.fetch(cal.room_id, cal.year, cal.week);
                }
            else{
                    cal.clearSelection()
                    cal.removeSettings();
                    cal.lb = true;

                    cal.fetch(cal.room_id, cal.year, cal.week);
                }
            }
        }
    },

    putNewBookingData : (e) => {
        //TODO

        const c = $(cal.r_c_o)
        c.innerText = "Occupancy: " + e[http.q.booking_occupancy]
        cal.cbkl(e)


    },

    /**
     * creates the query, for creating a new booking in the server
     * server has to check data again to accomodate for potential
     * injections or xss
     */
    createCreateQuery : () => {
        const y = cal.year;
        const w = cal.week;
        const m = cal.month;
        const d = cal.day;
        const start = cal.selected_column_min;
        const end = cal.selected_column_max
        const desc = $(cal.s_b_d).value
        const top = $(cal.s_b_t).value
        const count = $(cal.s_b_c).value


        //cal.



        const l = $(cal.rpbtn)
        if (l!=undefined) {
            if (cal.repeating){
                const e = $(cal.s_b_ed)
                const date = new Date(parseInt(e.value.split("-")[0]), parseInt(e.value.split("-")[1])-1, parseInt(e.value.split("-")[2]),0,0,0,0)
                const oneJan = new Date(date.getFullYear(),0,1);
                const ew = Math.ceil((date.valueOf() - oneJan.valueOf()) / (24 * 60 * 60 * 1000) / 7);


                return "?"+
                    http.q.booking_time_year + "=" + y + "&"+
                    http.q.booking_time_month + "=" + m + "&"+
                    http.q.booking_time_day + "=" + d + "&"+
                    http.q.booking_time_week + "=" + w + "&"+
                    http.q.booking_time_start + "=" + start + "&"+
                    http.q.booking_time_end + "=" + end + "&"+
                    http.q.booking_topic + "=" + top + "&"+
                    http.q.booking_description + "=" + desc + "&"+
                    http.q.booking_occupancy + "=" + count+ "&"+
                    http.q.booking_time_reoccuring + "=" + "true" + "&" +
                    http.q.booking_time_end_year + "=" + date.getFullYear() + "&"+
                    http.q.booking_time_end_month + "=" + (date.getMonth()+1) + "&"+
                    http.q.booking_time_end_day + "=" + (date.getDay()) + "&"+
                    http.q.booking_time_end_week + "=" + ew;
            }
        }
        return "?"+
            http.q.booking_time_year + "=" + y + "&"+
            http.q.booking_time_month + "=" + m + "&"+
            http.q.booking_time_day + "=" + d + "&"+
            http.q.booking_time_week + "=" + w + "&"+
            http.q.booking_time_start + "=" + start + "&"+
            http.q.booking_time_end + "=" + end + "&"+
            http.q.booking_topic + "=" + top + "&"+
            http.q.booking_description + "=" + desc + "&"+
            http.q.booking_occupancy + "=" + count + "&"+
            http.q.booking_time_reoccuring + "=" + "false"
    },

    /**
     * Clears the selection within the calender cell selection
     * to accomodate, when the week changes, the selection has been
     * submitted and a new booking has been created or on load
     */
    clearSelection : () => {
        const l = [];
        cal.selected_ids.forEach((e)=>{
            l.push(e);
        })
        for (let i=0;i<l.length;i++){
            const e = l[i]
            const k = parseInt(e.split("_")[1]);
            const j = parseInt(e.split("_")[0]);

            const o = $("m" + j + "_" + k)
            const d = cal.data["m"+j]["m"+j+"_"+k]

            cal.removeSelected(j,k);
            cal.cell(d,o);
        }
        cal.paintCurrentHour();
    },

    /**
     * Checks wether a cell from the currently selected column can be selected =>
     * Forces the selection to only consist of connected time cells from a single day
     * as well as the duration
     * @param j
     * @param k
     * @returns {boolean}
     */
    checkSelectedColumn : (j,k) => {
        if (!cal.auth)return false;
        if (cal.selected_ids.length >=cal.m_b_d)return false;
        if (cal.capacity<=cal.selected_ids.length + cal.r_o)return false;
        const e = (k>0)?(cal.data["m"+j]["m"+j+"_"+(k-1)].selected):(false)
        const l = (k<cal.nor-1)?(cal.data["m"+j]["m"+j+"_"+(k+1)].selected):(false)
        if (!cal.data["m"+j]["m"+j+"_"+(k)].bookable)return false;
        return e || l
    },

    /**
     * Returns the Week of the selected year from the calender as string
     * @returns {*|number}
     */
    getCalenderWeek : () => {
        return parseInt($(cal.cm).value.split("W")[1])
    },

    /**
     * Returns the Year of the selected calender as string
     * @returns {*|number}
     */
    getCalenderYear : () => {
        return parseInt($(cal.cm).value.split("-")[0])
    },

    getCalenderMonth : () => {
        let d = cal.weekDate(cal.year, cal.week)
        return d.getMonth();
    },

    /**
     * Removes the selected cell from the calender
     * @param j
     * @param k
     */
    removeSelected :(j,k) => {
        if (!cal.selected){
            return
        }
        var d = cal.data["m"+j]["m"+j+"_"+k]
        if (d.selected){
            const i = cal.selected_ids.indexOf(j + "_"+ k)
            cal.selected_ids.splice(i,1)
            cal.selected_occupancy = 0;
            cal.r_o = 0;
            cal.selected_ids.forEach((e) => {

                const i = e.charAt(0)
                const j = e.charAt(2)
                const p = cal.data["m"+i]["m"+i+"_"+j]
                if (p.occupancy >= cal.selected_occupancy){
                    cal.selected_occupancy = p.occupancy;
                    $(cal.r_c_o).innerText = "Occupancy : " + cal.selected_occupancy
                }
            });


            if (cal.selected_ids.length==0){
                cal.selected = false;
                $(cal.r_c_o).innerText = "Occupancy : -"
                const i = $(cal.c_b_btn);
                if (i!=undefined)i.setAttribute("disabled", true)

                const j = $(cal.rpbtn)
                if (j!=undefined){
                    j.style.backgroundColor = ""
                    j.classList.add("rpbtn", "ros-bg-sr")
                    j.setAttribute("disabled", true)
                }

                cal.repeating = false
                cal.deselectColumn(j)
                cal.paintCurrentHour()
            }
        if (cal.selected_column_min <k){
            cal.selected_column_min = k;
        }
        if (cal.selected_column_max >k){
            cal.selected_column_max = k;
        }
            d.selected = false;
        }
    },

    /**
     * Adds the current selected cell to the ids stack and
     * stores util values within the namespace
     * @param d
     * @param j
     * @param k
     */
    addSelected : (d, j,k) => {
        d.selected = true
        cal.selected_occupancy = d.occupancy;
        if (d.occupancy > cal.r_o)cal.r_o = d.occupancy;
        $(cal.s_b_c).max = (cal.capacity - cal.r_o);

        if (d.occupancy >= cal.selected_occupancy){
            $(cal.r_c_o).innerText = "Occupancy : " + cal.selected_occupancy
        }
        cal.selected_ids.push(j+"_"+k)
        if (cal.selected_column_min >k){
            cal.selected_column_min = k;
        }
        if (cal.selected_column_max <k){
            cal.selected_column_max = k;
        }
    },

    checkSelectedData : (d) => {
        cal.r_o = d.occupancy;
      return d.bookable &&  (cal.selected_ids.length < cal.m_b_d) && (cal.capacity>cal.selected_ids.length + cal.r_o) && cal.auth
    },

    /**
     * Checks wether the current selected cell from the calender can be removed
     * @param j
     * @param k
     * @returns {boolean}
     */
    checkRemoval : (j,k) => {
        let i=0;
        if ((k>0) && (cal.data["m"+j]["m"+j+"_"+(k-1)].selected)){
            i = i+1;
        }
        if ((k<cal.nor-1) && (cal.data["m"+j]["m"+j+"_"+(k+1)].selected)){
            i = i+1;
        }
        return i == 1 ||cal.selected_ids.length == 1
    },
}