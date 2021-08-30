/**
 *   main js file for implementing functionalities for the main menu
 *   this js file dynamicly fetches data from a common interface used on many websites
 *   that way the usermenubar, to select the rooms on each website can by
 *   integrated and reused as source code on multiple sides as microservice
 *
 *   @author ben
 */


const menu = {

    fetch : (resolve) => {
        const r = new XMLHttpRequest()
        r.open(http.me.fetch, "/api/menu", true)
        r.send()
        r.onreadystatechange = (e) => {
            if (r.readyState == XMLHttpRequest.DONE){
                const jsonData = JSON.parse(r.response);
                resolve(jsonData)
                //menu.createMenu(jsonData)
            }
        }
    },

    createMenu : (m0, j) => {
        const l = $$("li")
        $(m0).appendChild(l)
        const a = $$("a")
        a.innerText = "Rooms"
        a.setAttribute("href", "/")
        l.appendChild(a)
        const ul = $$("ul")

        ul.classList.add("m0ul")
        l.classList.add("m0l")
        a.classList.add("m0a")

        l.appendChild(ul);
        menu.cl1(j, ul)
    },

    dirs : ["N", "E", "S", "W"],

    cl1 : (j, u) => {
        for (let i=0;i<j["number_of_floors"];i++){
            const f = j["f"+i]
            const l = $$("li")
            u.appendChild(l)
            const a = $$("a")
            a.setAttribute("href", "#");
            a.innerText = f["floor"]
            l.appendChild(a)
            const ul = $$("ul")
            l.appendChild(ul)
            ul.id = "f"+i + "l0";

            ul.classList.add("m1ul")
            l.classList.add("m1l")
            a.classList.add("m1a")

            menu.dirs.forEach((e)=> {
                menu.cl2(f[e],ul, i, e);
            })
        }
    },

    cl2 : (j, u, i, e) => {
        const l = $$("li")
        l.classList.add("m2")
        u.appendChild(l)
        const a = $$("a")
        a.setAttribute("href", "#")
        a.innerText = j["direction"]
        l.appendChild(a)
        const ul = $$("ul")
        l.appendChild(ul)
        ul.id = "f"+i + "l1" + e  //d0l1W


        ul.classList.add("m2ul")
        l.classList.add("m2l")
        a.classList.add("m2a")

        j["rooms"].forEach((room, k)=>{
            menu.cl3(room, ul, i, e, k)
        })

    },

    cl3 : (r, u, i, e, k) => {
        const id = "f"+i + "l2" + e + k //f2l2W13 (etage 2, l2 => menu ebene räume, W = West, 13 raum
        const l = $$("li")
        u.appendChild(l)
        const a = $$("a")

        l.classList.add("m3l")
        a.classList.add("m3a")


        l.appendChild(a)
        a.setAttribute("href", "/rooms/"+r["room_id"]);
        a.innerText = r["room_name"];
    },

    init : (l) => {
        menu.fetch((e) =>{
            menu.createMenu(l,e)
        })
        //menu.createMenu(l, testdata.a)
    }
}