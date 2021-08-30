package controller

import (
	"encoding/json"
	"fmt"
	"mus_projekt/app/controller/protocol"
	"mus_projekt/app/controller/templates"
	"mus_projekt/app/model"
	"mus_projekt/app/model/query"
	"net/http"
	"strconv"
)

/**
	For all api templates, available to be stored serverside, altered through the webinterface and connected to a microcontroller
	this file contains all functionalities to respond to requsts and serve relwvant data

	@author ben
 */
func TemplateApi(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		html_templates.ExecuteTemplate(w, "display_template.tmpl", nil)
		return
	}

	if r.Method == protocol.MethodFetch {
		//create json data for the browser asnychronously

		id := r.FormValue(query.TEMPLATE_ID)
		fmt.Println(id)
		a, t, err := templates.CheckTemplateById(id)
		if !a || err != nil {
			m, err := templates.GetAllTemplates()
			if err != nil {
				m[query.SUCCESS] = false
			} else {
				m[query.SUCCESS] = true
			}
			b, _ := json.MarshalIndent(m, "", "  ")
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
			return
		}
		fmt.Println("HAHAHA")
		eid := r.FormValue(query.ELEMENT_ID)
		i, j := t.CheckElement(eid)
		if i {
			m := templates.CreateAPIData(t)
			m = m[query.TEMPLATE_ELEMENTS].([]map[string]interface{})[j]
			m[query.SUCCESS] = true
			b, _ := json.MarshalIndent(m, "", "  ")
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
			return
		} else {
			m := templates.CreateAPIData(t)
			fmt.Println(m)
			m[query.SUCCESS] = true
			b, _ := json.MarshalIndent(m, "", "  ")
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
			return
		}
	}
	if r.Method == http.MethodPost {
		//create json data for the browser asnychronously

		t := r.FormValue("t")
		if len(t) == 0 {

		} else {
			if t == "t" {
				//generate new template and return the json
				fmt.Println("Creating a new Template now")
				m := templates.CreateAPIData(templates.CreateNew())
				m[query.SUCCESS] = true
				b, _ := json.MarshalIndent(m, "", "  ")
				w.Header().Set("Content-Type", "application/json")
				w.Write(b)
				return
			}
			if t == "e" {
				id := r.FormValue(query.TEMPLATE_ID)
				fmt.Println(id)
				temp, _ := model.GetTemplateByTemplateID(id)
				e, _ := temp.AddElement()
				fmt.Println(">>>>>>>>>>>>>>>>>>>>The new element", e)
				m := templates.CreateAPIDataElement(temp.GetID(), e)
				//generate new element for the template id and return the json
				fmt.Println(m)
				fmt.Println(temp)
				m[query.SUCCESS] = true
				b, _ := json.MarshalIndent(m, "", "  ")
				w.Header().Set("Content-Type", "application/json")
				w.Write(b)
				return
			}
		}
	}
	if r.Method == http.MethodDelete {
		tid := r.FormValue(query.TEMPLATE_ID)
		eid := r.FormValue(query.ELEMENT_ID)
		fmt.Println("Remove the Data", tid, eid)
		a, t, err := templates.CheckTemplateById(tid)
		fmt.Println(a, t, err)
		if !a || err != nil {
			m := make(map[string]interface{})
			m[query.SUCCESS] = false
			b, _ := json.MarshalIndent(m, "", "  ")
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
			return
		}
		if a {
			fmt.Println("TEMPLATES EXISTS")
			if len(eid) == 0 {
				t.Remove()
				m, _ := templates.GetAllTemplates()
				m[query.SUCCESS] = true
				b, _ := json.MarshalIndent(m, "", "  ")
				w.Header().Set("Content-Type", "application/json")
				w.Write(b)
				return
			}
			i, j := t.CheckElement(eid)
			if i {
				t.RemoveElementByIndex(j)
				fmt.Println(t)
				m := templates.CreateAPIData(t)
				m[query.SUCCESS] = true
				b, _ := json.MarshalIndent(m, "", "  ")
				w.Header().Set("Content-Type", "application/json")
				w.Write(b)
				return
			}
		}
		m := make(map[string]interface{})
		m[query.SUCCESS] = false
		b, _ := json.MarshalIndent(m, "", "  ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		return
	}

	if r.Method == http.MethodPut {
		fmt.Println("\n\nSAVE NOW")
		tid := r.FormValue(query.TEMPLATE_ID)
		eid := r.FormValue(query.ELEMENT_ID)
		fmt.Println(tid, eid)
		a, t, err := templates.CheckTemplateById(tid)
		if !a || err != nil {
			m := make(map[string]interface{})
			m[query.SUCCESS] = false
			b, _ := json.MarshalIndent(m, "", "  ")
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
			return
		} else {
			i, j := t.CheckElement(eid)
			if !i {
				tnm := r.FormValue(query.TEMPLATE_NAME)
				t.SetName(tnm)
				m := make(map[string]interface{})
				m[query.SUCCESS] = true
				q, _ := json.MarshalIndent(m, "", "  ")
				w.Header().Set("Content-Type", "application/json")
				w.Write(q)
				return
			} else {
				fmt.Println("HEHEHEHEHEHEHEHEHEHEH")
				b, _ := strconv.ParseBool(r.FormValue(query.ELEMENT_CONTENT_STATIC))
				x_, _ := strconv.Atoi(r.FormValue(query.ELEMENT_POSITION_X))
				y_, _ := strconv.Atoi(r.FormValue(query.ELEMENT_POSITION_Y))
				w_, _ := strconv.Atoi(r.FormValue(query.ELEMENT_POSITION_W))
				h_, _ := strconv.Atoi(r.FormValue(query.ELEMENT_POSITION_H))
				fs, _ := strconv.Atoi(r.FormValue(query.ELEMENT_FONT_SIZE))
				ps, _ := strconv.Atoi(r.FormValue(query.ELEMENT_PIXEL_SIZE))

				fmt.Println(x_, y_, w_, h_)
				t.SaveSettings(
					j,
					r.FormValue(query.TEMPLATE_NAME),
					r.FormValue(query.ELEMENT_ID),
					r.FormValue(query.ELEMENT_CONTENT),
					b,
					x_,
					y_,
					w_,
					h_,
					r.FormValue(query.ELEMENT_COLOR),
					r.FormValue(query.ELEMENT_FILL_COLOR),
					fs,
					ps,
					r.FormValue(query.ELEMENT_PIXEL_STYLE),
					r.FormValue(query.ELEMENT_FORM))

				m := make(map[string]interface{})
				m[query.SUCCESS] = true
				q, _ := json.MarshalIndent(m, "", "  ")
				w.Header().Set("Content-Type", "application/json")
				w.Write(q)
				return
			}
		}
	}

}
