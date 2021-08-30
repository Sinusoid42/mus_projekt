/**
 * User Page functionalities
 * @type {{open: u.open}}
 *
 * @author ben
 */




const u = {

    search : "si",  //search input

    //Searches for a name in the database
    find : ()=> {
        const r = new XMLHttpRequest()
        r.open(http.me.find, "http://" + http.a.address + ":" + http.a.port + http.p.users + "?find=" + $("si").value)

    },
    open : (uri) => {
        if (http.p.login.includes(uri)){
            window.open("http://" + http.a.address + ":" + http.a.port + http.p.login, "_self")
        }
        if (http.p.logout.includes(uri)){
            window.open("http://" + http.a.address + ":" + http.a.port + http.p.logout, "_self")
        }
        if (http.p.bookings.includes(uri)){
            window.open("http://" + http.a.address + ":" + http.a.port + http.p.bookings, "_self")
        }
        if (http.p.users.includes(uri)){
            window.open("http://" + http.a.address + ":" + http.a.port + http.p.users, "_self")
        }
        if (uri.includes("api") && uri.includes("rooms")){
            window.open("http://" + http.a.address + ":" + http.a.port + http.p.room_api, "_self")
        }

        if (uri.includes("api") && uri.includes("templates")){
            window.open("http://" + http.a.address + ":" + http.a.port + http.p.template_api, "_self")
        }
        if (uri.includes("api") && uri.includes("arduino") ){
            window.open("http://" + http.a.address + ":" + http.a.port + http.p.arduino_api, "_self")
        }
        if (uri.includes("api") && uri.includes("upgrades") ){
            window.open("http://" + http.a.address + ":" + http.a.port + http.p.upgrade_api, "_self")
        }
    },

    initCAL : () => {
        if (cal && cal != undefined) {
            cal.init();
        }
    },
}




window.onload = () => {
    api.resolveServerAddress.then((server_data_json) => {
        http.putServerData(server_data_json);
        menu.init("m0")

        new Promise((s) => {
            if (cal != undefined){
                s(cal)
            }
        }).then((e) => {
            e.init()
        })
        new Promise((s) => {
            if (acc != undefined){
                s(acc)
            }
        }).then((e) => {
            e.init()
        })

    })
}


































