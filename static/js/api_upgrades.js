/**
 *
 * Main JS file to access the api page when upgrading accounts
 * This file sends asyncrhonous data to the server which in turn sends Emails to the specified user account
 *
 * @type {{container: string, init: uapi.init, createRequestField: uapi.createRequestField, fetch: uapi.fetch, g: uapi.g, save: uapi.save}}
 *
 * @author ben
 */



const uapi = {

    container : "upgrade_container",

    fetch : () => {
        $(uapi.container).innerHTML=""
        const r = new XMLHttpRequest();
        r.open(http.me.fetch, http.p.upgrade_api, true);
        r.send();
        r.onreadystatechange = () => {
            if (r.readyState == XMLHttpRequest.DONE){
                const j = JSON.parse(r.response)
                console.log(j);
                if (j[http.q.success]){
                    j["upgrades"].forEach((e)=> {
                        uapi.createRequestField(e);
                    })
                }
            }
        }
    },
    g : (i) => {
        uapi.v = $("ip_"+i).value
    },

    //when saving the ticket is not destroyed
    save : (id) => {
        uapi.g(id)
        const r = new XMLHttpRequest();
        r.open(http.me.post, http.p.upgrade_api  + "?id=" + id + "&t=t0" + "&l=" + uapi.v, true);
        r.send();
        r.onreadystatechange = () => {
            if (r.readyState == XMLHttpRequest.DONE){
                const j = JSON.parse(r.response)
                console.log(j);

            }
        }
    },

    init : () => {
        uapi.fetch();
    },

    createRequestField : (j) => {
        console.log(j)
        const e = document.createElement("div")
        e.classList.add("row", "mt-2", "justify-content-around")
        const s = ` 
                        <div class="col-4 font-weight-bold text-center">
                                                Email Address
                        </div>
                        <div class="col-4 font-weight-bold text-center">
                                               `+ j["user_email_address"]+`
                        </div>
                        <div class="col-2 font-weight-bold">
                                <div class="row justify-content-around">
                                        <input class="col-6" value="`+ j["admin_access_level"] +`"  min="0" max="5" type="number" id ="ip_`+ j["upgrade_helper_id"] +`">
                                    
                                    <div class="col-6 font-weight-bold text-center upgrade_confirm" onclick="uapi.save('`+ j["upgrade_helper_id"] + `')">
                                        Send
                                    </div>
                                </div>
                        </div>
                   `;
        e.innerHTML = s
        $(uapi.container).appendChild(e)
    },
}


window.onload = () => {
    uapi.init();
}