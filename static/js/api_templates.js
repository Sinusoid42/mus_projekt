/*
    Main js implementation for the display_template.tmpl

 */
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

 Templates api website

 this frontend webpage delivers the functionalities to create and remove
 templates from the renderable templates content stored within the database

 when a user has the access given though his account, otherwise rejected
 he will have access to all templates rendered on the arduinos

 the arduino has to be connected via

 @author ben

 */
const settings = {
    template_id : "",
    template_name : "",

    element_id              : "",
    element_content         : "",
    element_content_static  : false,
    element_x               : 0,
    element_y               : 0,
    element_w               : 0,
    element_h               : 0,
    element_color           : "",
    element_fill_color      : "",
    element_font_size       : 0,
    element_pixel_size      : 0,
    element_style           : "",
    element_form            : "",
}

const t = {
    tc : "template_container",
    ec : "element_container",
    atb : "atb",
    aeb : "aeb",

    st : "",
    se : "",

    seE : null,
    stE : null,

    settings_t_id : "t_tid",
    settings_t_name : "t_tnm",
    settings_e_id : "t_eid",

    settings_e_dynamic_content : "t_edc",
    settings_e_content_static : "t_cst",
    settings_e_x : "t_ex",
    settings_e_y : "t_ey",
    settings_e_w : "t_ew",
    settings_e_h : "t_eh",
    settings_e_color : "t_ec",
    settings_e_f_color : "t_efc",
    settings_e_f_size : "t_fs",
    settings_e_p_size : "t_psi",
    settings_e_p_style : "t_pst",
    settings_e_form : "t_frm",
    settings_e_s_btn : "t_btn",

    list_template_name : "tnm",

    /**
     * Data Communication Setup and Callbacks
     */
    //retrieves all Data from the Server to Display all Templates => check paths in go server
    fetch : () => {
        //TODO Fetch all the templates from the database and send it to the client
        const r  = new XMLHttpRequest();
        r.open(http.me.fetch, http.p.template_api, true);
        r.send();
        r.onreadystatechange = (e) => {
            if (r.readyState == XMLHttpRequest.DONE){
                const j = JSON.parse(r.response);
                if (j[http.q.success]){
                    t.putTemplates(j)
                }
            }
        }
    },

    ct : () => {
        const r  = new XMLHttpRequest()
        r.open(http.me.post, http.p.template_api + "?t=t", true)
        r.send()
        r.onreadystatechange = (e) => {
            if (r.readyState == XMLHttpRequest.DONE){
                const j = JSON.parse(r.response)
                if (j["success"]){
                    t.createTemplate(j);
                }
            }
        }
    },

    ce : () => {
        if (t.stE == null)return;
        const r  = new XMLHttpRequest()
        r.open(http.me.post, http.p.template_api + "?t=e" + "&"+ http.q.template_id +"="+t.st, true)
        r.send()
        r.onreadystatechange = (e) => {
            if (r.readyState == XMLHttpRequest.DONE){
                const j = JSON.parse(r.response)
                if (j[http.q.success]){
                    t.createElement(j, j[http.q.template_id])
                }
            }
        }
    },

    fetchT : (id, cb) => {
        const r = new XMLHttpRequest();
        r.open(http.me.fetch, http.p.template_api + "?" + http.q.template_id + "=" + id, true)
        r.send()
        r.onreadystatechange = (e) => {
            if (r.readyState == XMLHttpRequest.DONE) {
                const j = JSON.parse(r.response)
                if (j[http.q.success]){
                    cb(j)
                }
            }
        }
    },

    fetchE : (tid, eid, cb) => {
        const r = new XMLHttpRequest();
        r.open(http.me.fetch, http.p.template_api + "?" + http.q.template_id + "=" + tid + "&" + http.q.template_element_id + "="  + eid, true)
        r.send()
        r.onreadystatechange = (e) => {
            if (r.readyState == XMLHttpRequest.DONE) {
                const j = JSON.parse(r.response)
                if (j[http.q.success]){
                    cb(j)
                }
            }
        }
    },

    /**
     * Data Communication Event Handling
     */

    clearElements: () => {
      $(t.ec).innerHTML = ''
    },

    clearTemplates : () => {
        $(t.tc).innerHTML = ''
    },

    ceml : (e) => {
        return '<div class="row mt-2"><div class="col-12 text-center " ' +
            'style="font-size: 12px;"><b> Position </b></div></div><div ' +
            'class="row justify-content-center">' +
            '<div class="col-2 text-center text-white ros-bg-db"style="border-style: solid;border-color: #000000;border-radius: 5px" id="'+ e[http.q.template_element_id] + '_x">' + e[http.q.template_element_position_x]+'</div>' +
            '<div class="col-2 text-center text-white ros-bg-db"style="border-style: solid;border-color: #000000;border-radius: 5px" id="'+ e[http.q.template_element_id] + '_y">' + e[http.q.template_element_position_y]+'</div>' +
            '<div class="col-2 text-center text-white ros-bg-db"style="border-style: solid;border-color: #000000;border-radius: 5px" id="'+ e[http.q.template_element_id] + '_w">' + e[http.q.template_element_position_w]+'</div>' +
            '<div class="col-2 text-center text-white ros-bg-db"style="border-style: solid;border-color: #000000;border-radius: 5px" id="'+ e[http.q.template_element_id] + '_h">' + e[http.q.template_element_position_h]+'</div>' +
            '</div>'
    },

    cel : (i, e , n, h) => {
        return '<div class="row mt-2"><div class="col-12 text-center " ' +
            'style="font-size: 12px;"><b>' + h + '</b></div></div><div ' +
            'class="row justify-content-center"><div class="col-12 text-center text-white ros-bg-db" ' +
            'style="border-style: solid;border-color: #000000;border-radius: 5px" id="'+ i + '">' + e[n] + '</div></div>'
    },

    selectT : (e, id) => {
        if (t.stE != null){
           t.stE.style.backgroundColor = color_sceme.sb;
        }
        if (t.stE != e){
            if (t.seE != null){
                t.seE.style.backgroundColor = color_sceme.sb;
            }
            t.seE = null;
            t.emptyElementSettings();
        }

        e.style.backgroundColor = color_sceme.dr;
        t.stE = e;
        t.st = id;
        t.fetchT(id , (e) => {
            t.putTemplateSettings(e)
          t.putElements(e)
        });
    },

    //Selects an Element from a Template
    selectE : (e, tid, eid) => {
        if (t.seE != null){
            t.seE.style.backgroundColor = color_sceme.sb;
        }
        e.style.backgroundColor = color_sceme.dr;
        t.seE = e;
        t.se = eid;

        t.fetchE(tid, eid, (e)=> {
            t.putElementSettings(e)
        })
    },

    emptyElementSettings : () => {
        $(t.settings_t_id).innerText = settings.template_id;
        $(t.settings_t_name).value = settings.template_name;
        $(t.settings_e_id).innerText = "-";
        $(t.settings_e_dynamic_content).value = "-";
        //$(t.settings_t_id).innerText = settings.template_id; //TODO

        $(t.settings_e_content_static).style.backgroundColor = color_sceme.lr;

        $(t.settings_e_x).value = 0;
        $(t.settings_e_y).value = 0;
        $(t.settings_e_w).value = 0;
        $(t.settings_e_h).value = 0;

        $(t.settings_e_color).value = "-";
        $(t.settings_e_f_color).value = "-";
        $(t.settings_e_f_size).value = 0;
        $(t.settings_e_p_size).value = 0;
        $(t.settings_e_p_style).value = "-";
        $(t.settings_e_form).value = "-";
    },

    emptySettings : () => {
        $(t.settings_t_id).innerText = "-";
        $(t.settings_t_name).value = "-";
        $(t.settings_e_id).innerText = "-";
        $(t.settings_e_dynamic_content).value = "-";
        //$(t.settings_t_id).innerText = settings.template_id; //TODO

        $(t.settings_e_content_static).style.backgroundColor = color_sceme.lr;

        $(t.settings_e_x).value = 0;
        $(t.settings_e_y).value = 0;
        $(t.settings_e_w).value = 0;
        $(t.settings_e_h).value = 0;

        $(t.settings_e_color).value = "-";
        $(t.settings_e_f_color).value = "-";
        $(t.settings_e_f_size).value = 0;
        $(t.settings_e_p_size).value = 0;
        $(t.settings_e_p_style).value = "-";
        $(t.settings_e_form).value = "-";
    },

    putElements : (e) => {
        t.clearElements();
        e["elements"].forEach((l, k) => {
            t.createElement(l, e[http.q.template_id])
        })
    },

    putTemplates : (e) => {
        t.clearElements();
        t.clearTemplates();

        e["templates"].forEach((l, i) => {
            t.createTemplate(l)
        })
    },

    /**
     * User Input Event Handling
     */
    //removes a Template
    removeT : (tid) => {
        const r = new XMLHttpRequest()
        r.open(http.me.delete, http.p.template_api + "?" + http.q.template_id + "=" + tid, true)
        r.send()
        r.onreadystatechange  = () => {
            if (r.readyState == XMLHttpRequest.DONE){
                const j = JSON.parse(r.response)
                if (j[http.q.success]){
                    t.emptySettings();
                    t.putTemplates(j);
                    //TODO REMOVE THE ROOM AND THE ELEMENTS WHEN SELECTED => reset settings to default
                }
            }
        }
    },

    //removes a Element from a Template
    removeE : (tid, eid) => {
        const r = new XMLHttpRequest()
        r.open(http.me.delete, http.p.template_api + "?" + http.q.template_id + "=" + tid + "&" + http.q.template_element_id + "=" + eid, true)
        r.send()
        r.onreadystatechange  = () => {
            if (r.readyState == XMLHttpRequest.DONE){
                const j = JSON.parse(r.response)
                if (j[http.q.success]){
                    t.putElements(j)
                    t.emptyElementSettings()
                    //TODO REMOVE ELEMENT WHEN SELECTED => remove settings elements only
                }
            }
        }
    },

    //initializes the js implementations for the template for creating arduino templates to be rendererd
    init : () => {
        $(t.settings_e_s_btn).onclick = () => {
            //TODO SAVE METHOD
            t.save()

        }
        /**
         * Add Template Button Onclick
         */
        $(t.atb).onclick = () => {
            t.ct(); //create template

        }
        /**
         * Add Element Button Onclick
         */
        $(t.aeb).onclick = () => {
            t.ce(); //create template
        }

        $(t.settings_e_content_static).onclick = () => {
            if (t.seE == null){
                return
            }
            console.log(t.seE)
            settings.element_content_static = !settings.element_content_static
            if ( settings.element_content_static){
                $(t.settings_e_content_static).style.backgroundColor = '#80F530';
            }
            else{
                $(t.settings_e_content_static).style.backgroundColor = color_sceme.lr;
            }

        }

        t.fetch()//checks for all templates in the database and puts them into the list, with the first template with elements into the element list

    },

    putTemplateSettings : (e) => {
        settings.template_id = e[http.q.template_id]
        settings.template_name = e[http.q.template_name]
        $(t.settings_t_id).innerText = settings.template_id;
        $(t.settings_t_name).value = settings.template_name;
    },

    setElement : () => {
        try {
            console.log(settings.element_id + t.settings_e_dynamic_content)
            $(settings.element_id + t.settings_e_dynamic_content).innerHTML = settings.element_content
            $(settings.element_id + t.settings_e_color).innerHTML = settings.element_color
            $(settings.element_id + t.settings_e_f_color).innerHTML = settings.element_fill_color

            $(settings.element_id + "_x").innerHTML = settings.element_x
            $(settings.element_id + "_y").innerHTML = settings.element_y
            $(settings.element_id + "_w").innerHTML = settings.element_w
            $(settings.element_id + "_h").innerHTML = settings.element_h
        }
        catch(e){}
    },

    setTemplate : () => {
        $(settings.template_id + t.settings_t_name).innerHTML = settings.template_name
    },

    putElementSettings : (e) => {
        settings.element_id = e[http.q.template_element_id]
        settings.element_content = e[http.q.template_element_content]
        settings.element_content_static = e[http.q.template_element_content_static]
        settings.element_x = e[http.q.template_element_position_x]
        settings.element_y = e[http.q.template_element_position_y]
        settings.element_w = e[http.q.template_element_position_w]
        settings.element_h = e[http.q.template_element_position_h]
        settings.element_color = e[http.q.template_element_color]
        settings.element_fill_color = e[http.q.template_element_fill_color]
        settings.element_style = e[http.q.template_element_style]
        settings.element_form = e[http.q.template_element_form]
        settings.element_font_size = e[http.q.template_element_font_size]
        settings.element_pixel_size = e[http.q.template_element_pixel_size]

        $(t.settings_e_id).innerText = settings.element_id;
        $(t.settings_e_dynamic_content).value = settings.element_content;
        //$(t.settings_t_id).innerText = settings.template_id; //TODO
        if (settings.element_content_static){
            $(t.settings_e_content_static).style.backgroundColor = '#80F530';
        }else{
            $(t.settings_e_content_static).style.backgroundColor = color_sceme.lr;
        }
        $(t.settings_e_x).value = settings.element_x;
        $(t.settings_e_y).value = settings.element_y;
        $(t.settings_e_w).value = settings.element_w;
        $(t.settings_e_h).value = settings.element_h;

        $(t.settings_e_color).value = settings.element_color;
        $(t.settings_e_f_color).value = settings.element_fill_color;
        $(t.settings_e_f_size).value = settings.element_font_size;
        $(t.settings_e_p_size).value = settings.element_pixel_size;
        $(t.settings_e_p_style).value = settings.element_style;
        $(t.settings_e_form).value = settings.element_form;
    },

    loadFromSettingsForm : () => {
        settings.template_id = $(t.settings_t_id).innerText
        settings.template_name = $(t.settings_t_name).value
        settings.element_id = $(t.settings_e_id).innerText
        settings.element_content = $(t.settings_e_dynamic_content).value
        //settings.element_content_static //TODO
        settings.element_x = $(t.settings_e_x).value
        settings.element_y = $(t.settings_e_y).value
        settings.element_w = $(t.settings_e_w).value
        settings.element_h = $(t.settings_e_h).value
        settings.element_color = $(t.settings_e_color).value
        settings.element_fill_color = $(t.settings_e_f_color).value
        settings.element_pixel_size =  $(t.settings_e_p_size).value
        settings.element_font_size= $(t.settings_e_f_size).value
        settings.element_style = $(t.settings_e_p_style).value
        settings.element_form = $(t.settings_e_form).value

    },

    //puts the settings into the ui menu from inner storage
    putSettings : () => {
        $(t.settings_t_id).innerText = settings.template_id;
        $(t.settings_t_name).value = settings.template_name;
        $(t.settings_e_id).innerText = settings.element_id;
        $(t.settings_e_dynamic_content).value = settings.element_content;
        //$(t.settings_t_id).innerText = settings.template_id; //TODO
        if (settings.element_content_static){
            $(t.settings_e_content_static).style.backgroundColor = '#80F530';
        }else{
            $(t.settings_e_content_static).style.backgroundColor = color_sceme.lr;
        }
        $(t.settings_e_x).value = settings.element_x;
        $(t.settings_e_y).value = settings.element_y;
        $(t.settings_e_w).value = settings.element_w;
        $(t.settings_e_h).value = settings.element_h;

        $(t.settings_e_color).value = settings.element_color;
        $(t.settings_e_f_color).value = settings.element_fill_color;
        $(t.settings_e_f_size).value = settings.element_font_size;
        $(t.settings_e_p_size).value = settings.element_pixel_size;
        $(t.settings_e_p_style).value = settings.element_style;
        $(t.settings_e_form).value = settings.element_form;

    },

    //saves the contents of the settings available for the template element into the inner storage => cb store in db via server
    save : () => {
        const r = new XMLHttpRequest();
        t.loadFromSettingsForm()
        r.open(http.me.put, t.createQuery(), true)
        r.send()
        r.onreadystatechange = (e) => {
            if (r.readyState == XMLHttpRequest.DONE){
                const j = JSON.parse(r.response);
                if (j[http.q.success]){
                    t.setElement();
                    t.setTemplate();
                }
            }
        }
    },

    createQuery : () => {
        return "?"+
            http.q.template_id + "=" + settings.template_id + "&"+
            http.q.template_name + "=" + settings.template_name +"&"+
            http.q.template_element_id+ "=" + settings.element_id +"&"+
            http.q.template_element_content + "=" + settings.element_content +"&"+
            http.q.template_element_content_static + "=" + settings.element_content_static +"&"+
            http.q.template_element_position_x+ "=" + settings.element_x +"&"+
            http.q.template_element_position_y+ "=" + settings.element_y +"&"+
            http.q.template_element_position_w+ "=" + settings.element_w +"&"+
            http.q.template_element_position_h+ "=" + settings.element_h +"&"+
            http.q.template_element_color + "=" + settings.element_color +"&"+
            http.q.template_element_fill_color + "=" + settings.element_fill_color +"&"+
            http.q.template_element_font_size + "=" + settings.element_font_size +"&"+
            http.q.template_element_pixel_size + "=" + settings.element_pixel_size +"&"+
            http.q.template_element_style + "=" + settings.element_style +"&"+
            http.q.template_element_form + "=" + settings.element_form
    },

    //creates a new Template HTML Element
    createTemplate : (e) => {
        const r = $$("row")
        r.classList.add("justify-content-center")
        r.id = e[http.q.template_id]
        r.innerHTML = '<div class="col-12 ros-bg-sb mt-2 mb-3 text-center"' +
            'onclick="t.selectT(this, `'+ e[http.q.template_id] +'`)" ' +
            'style="border-style: solid;' +
            'border-radius: 1em"><div class="row justify-content-center"><div class="col-11 mt-2 mb-2">' +
            '<div class="row justify-content-end"><div class="col-9" id="'+e[http.q.template_id]+t.settings_t_name+'"><b><u>' + e[http.q.template_name] +
            '</u></b></div><button onclick="t.removeT(`' + e[http.q.template_id]  + '`)" " class="col-3 btn ros-bg-sr text-center" ' +
            'style="width: 2em;height: 2em;font-size: 12px;padding-bottom: 1.6em;"><b>X</b></button></div>'+
            t.cel(e[http.q.template_id]+t.settings_t_id, e, http.q.template_id, "TemplateID") +
            '</div></div></div>'
        $(t.tc).appendChild(r)
    },

    //creates a new Element HTML Element
    createElement : (e, tid) => {
        console.log(e[http.q.template_element_id]+t.settings_e_dynamic_content)
        const r = $$("row")
        r.classList.add("justify-content-center")
        r.id = e[http.q.template_element_id]
        r.innerHTML = '<div class="col-12 ros-bg-sb mt-2 mb-3 text-center"' +
            'onclick="t.selectE(this,`'+ tid +'`, `'+ e[http.q.template_element_id] +'`)" ' +
            'style="border-style: solid;' +
            'border-radius: 1em"><div class="row justify-content-center"><div class="col-11 mt-2 mb-2">' +
            '<div class="row justify-content-end"><button onclick="t.removeE(`'+ tid + '`,`' + e[http.q.template_element_id]  + '`)" " class="col-3 btn ros-bg-sr text-center" ' +
            'style="width: 2em;height: 2em;font-size: 12px;padding-bottom: 1.6em;"><b>X</b></button></div>'+
            t.cel(e[http.q.template_element_id]+t.settings_e_id, e, http.q.template_element_id, "ElementID") +
            t.cel(e[http.q.template_element_id]+t.settings_e_dynamic_content, e, http.q.template_element_content, "Content") +
            t.cel(e[http.q.template_element_id]+t.settings_e_color, e, http.q.template_element_color, "Color") +
            t.cel(e[http.q.template_element_id]+t.settings_e_f_color, e, http.q.template_element_fill_color, "Fill Color") +
            t.ceml(e) +
            '</div></div></div>'
        $(t.ec).appendChild(r)
    },
}

window.onload = () => {
    api.resolveServerAddress.then((e)=> {
        http.putServerData(e);
        t.init()
    })
}




