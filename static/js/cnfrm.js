/**
 *
 *   js implementations for the account confirmation site
 *
 * @author ben
 */

const submitbtn = "submitbtn";
let d;
const u = {
    em:"",
    update : () => {
        const d = $(submitbtn)
        d.onclick = () => {
            const e0 = $("u")
            const e1 = $("p0")
            const e2 = $("p1")
            const b0 = $("submitbtn")


            if (e1.value != e2.value || e1.value.length == 0 || e2.value.length == 0 || e0.value.length == 0){

                lock = true
                e1.style.color = color_sceme.dr
                e2.style.color = color_sceme.dr
                b0.style.backgroundColor = color_sceme.dr
                setTimeout(() => {
                    lock = false
                    e1.style.color = "#000000"
                    e2.style.color = "#000000"
                    b0.style.backgroundColor = color_sceme.lr
                }, 1500);
                return
            }

            const r = new XMLHttpRequest()
            const addr = http.p.confirmation + "?"+ http.q.user_email + "=" + u.em + "&"+http.q.user_name +"=" + e0.value + "&" + http.q.user_password + "=" + encodeRot13(e1.value)
            console.log("The address")
            console.log(addr)
            r.open(http.me.post, addr, true)
            r.send()
            r.onreadystatechange = () => {
                if (r.readyState == XMLHttpRequest.DONE){
                    const j = JSON.parse(r.response)
                    if (j[http.q.success]){

                        $("a0").innerText ="Success, redirecting to login. . . "
                        setTimeout(() => {
                            const addr ="http://" + http.a.address + ":" + http.a.port + http.p.login
                            console.log(addr)
                            window.open(addr, "_self")
                        },1000)
                    }


                }
            }
        }
    }
}
let id = window.location.search.substr(4);

let lock = false;
window.onload =() => {
    api.resolveServerAddress.then((e)=> {
        http.putServerData(e)
    })
     d = document.getElementById(submitbtn)
    d.onmouseenter = () => {
         if (lock)return
         d.style.backgroundColor = color_sceme.sr
    }
    d.onmouseleave = () => {
        if (lock)return
        d.style.backgroundColor = color_sceme.lr
    }

}





