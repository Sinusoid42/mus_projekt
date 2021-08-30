/**
 *   Clientside implementations to create functionalities on the
 *   password reset website
 *
 *   this communication handles the saving of a new entered password
 *
 *   @author ben
 */

let lock = false;
const a = {
      f : () => {
          window.location.href.split()
          const i0 = $("i0");
          const i1 = $("i1");
          if (!lock) {
              lock = true;
              if (i0.value != i1.value || i0.value.length == 0 || i1.value.length == 0) {
                  $("b0").style.backgroundColor = color_sceme.pr
                  setTimeout(() => {
                      $("b0").style.backgroundColor = color_sceme.lr
                      lock = false
                  }, 800);
              } else {
                  const id = window.location.href.split("=")[1]//Is the id for the page
                  $("b0").style.backgroundColor = "#A0F0A0"
                  setTimeout(() => {
                      $("b0").style.backgroundColor = color_sceme.lr
                      lock = false
                  }, 800);
                  const r = new XMLHttpRequest();
                  r.open(http.me.post, http.p.password_forget + "?id=" + id + "&" + http.q.user_password + "=" + encodeRot13(i0.value))
                  r.send();
                  r.onreadystatechange = () => {
                      if (r.readyState == XMLHttpRequest.DONE) {
                          console.log(r.response)
                          const j = JSON.parse(r.response);
                          console.log(j)
                          if (j["success"]){
                                window.open("http://" + http.a.address + ":" + http.a.port + http.p.login, "_self")
                          }
                      }
                  }
              }
          }
      },
}

window.onload = () => {
    api.resolveServerAddress.then((e)=> {
        http.putServerData(e);
    })
}