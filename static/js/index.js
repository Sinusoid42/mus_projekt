/**
 *
 * Javascript file to possibly extend the functionalities of the main
 * front page when using the room occupancy system
 *
 *
 * @type {{M0: string, m0: string, mbtn1: string, mbtn0: string}}
 *
 * @author ben
 */

const index = {
    mbtn0 : "M0btn0",
    mbtn1 : "M0btn1",
    M0 : "M0",
    m0 : "m0"

}

window.onload = () => {
    api.resolveServerAddress.then((e)=> {
        http.putServerData(e)
    })
    menu.init(index.m0)


    $(index.mbtn0)
    console.log("AYE")


}