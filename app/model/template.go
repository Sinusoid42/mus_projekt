package model

import (
	"errors"
	"fmt"
	"log"
	"mus_projekt/app/model/query"
	"mus_projekt/app/model/utils"
)

type Template struct {
	_id  string `json:"_id"`
	_rev string `json:"_rev"`

	template_id   string     `json:"template_id"`
	template_name string     `json:"template_name"`
	elements      []*Element `json:"elements"`
}

type Element struct {
	element_id     string `json:"element_id"`
	content        string `json:"content"`
	content_static bool   `json:"content_static"`
	x              int    `json:"x"`
	y              int    `json:"y"`
	w              int    `json:"w"`
	h              int    `json:"h"`
	color          string `json:"color"`
	fill_color     string `json:"fill_color"`
	font_size      int    `json:"font_size"`
	pixel_size     int    `json:"pixel_size"`
	style          string `json:"style"`
	form           string `json:"form"`
}

func t2m(t *Template) map[string]interface{} {
	m := make(map[string]interface{})
	m[query.DB_ID] = t._id
	m[query.DB_REV] = t._rev
	m[query.TEMPLATE_ID] = t.template_id
	m[query.TEMPLATE_NAME] = t.template_name
	m[query.TEMPLATE_ELEMENTS] = parse_template_elements_e2m(t.elements)
	return m
}

func parse_template_elements_e2m(e []*Element) []map[string]interface{} {
	m := []map[string]interface{}{}
	for _, k := range e {
		m = append(m, e2m(k))
	}
	return m
}

func parse_template_elements_m2e(m interface{}) ([]*Element, error) {
	e := []*Element{}
	for _, k := range m.([]interface{}) {
		j, err := m2e(k.(map[string]interface{}))
		if err != nil {
			return nil, err
		}
		e = append(e, j)
	}
	return e, nil
}

func e2m(e *Element) map[string]interface{} {
	j := make(map[string]interface{})

	j[query.ELEMENT_ID] = e.element_id
	j[query.ELEMENT_CONTENT] = e.content
	j[query.ELEMENT_CONTENT_STATIC] = e.content_static
	j[query.ELEMENT_POSITION_X] = e.x
	j[query.ELEMENT_POSITION_Y] = e.y
	j[query.ELEMENT_POSITION_W] = e.w
	j[query.ELEMENT_POSITION_H] = e.h
	j[query.ELEMENT_COLOR] = e.color
	j[query.ELEMENT_FILL_COLOR] = e.fill_color
	j[query.ELEMENT_FONT_SIZE] = e.font_size
	j[query.ELEMENT_PIXEL_SIZE] = e.pixel_size
	j[query.ELEMENT_PIXEL_STYLE] = e.style
	j[query.ELEMENT_FORM] = e.form
	return j
}

func m2e(m map[string]interface{}) (*Element, error) {
	if m == nil {
		return nil, errors.New("The element Map could not be found")
	}
	e := Element{
		element_id:     m[query.ELEMENT_ID].(string),
		content:        m[query.ELEMENT_CONTENT].(string),
		content_static: m[query.ELEMENT_CONTENT_STATIC].(bool),
		x:              int(m[query.ELEMENT_POSITION_X].(float64)),
		y:              int(m[query.ELEMENT_POSITION_Y].(float64)),
		w:              int(m[query.ELEMENT_POSITION_W].(float64)),
		h:              int(m[query.ELEMENT_POSITION_H].(float64)),
		color:          m[query.ELEMENT_COLOR].(string),
		fill_color:     m[query.ELEMENT_FILL_COLOR].(string),
		font_size:      int(m[query.ELEMENT_FONT_SIZE].(float64)),
		pixel_size:     int(m[query.ELEMENT_PIXEL_SIZE].(float64)),
		style:          m[query.ELEMENT_PIXEL_STYLE].(string),
		form:           m[query.ELEMENT_FORM].(string),
	}
	return &e, nil
}

func m2t(m map[string]interface{}) (*Template, error) {

	e, err := parse_template_elements_m2e(m[query.TEMPLATE_ELEMENTS])
	if err != nil {
		return nil, err
	}
	t := Template{
		_id:           m[query.DB_ID].(string),
		_rev:          m[query.DB_REV].(string),
		template_id:   m[query.TEMPLATE_ID].(string),
		template_name: m[query.TEMPLATE_NAME].(string),
		elements:      e,
	}
	return &t, nil

}

func (t *Template) GetElementByID(id string) (*Element, error) {

	for _, e := range t.elements {
		if e.element_id == id {
			return e, nil
		}
	}
	return nil, errors.New("The Element could not be found")

}

func GetTemplateByTemplateID(id string) (*Template, error) {
	query := `{"selector" : {"template_id" : "` + id + `" }}`
	m, err := template_DB.QueryJSON(query)
	if err != nil {
		return nil, err
	}
	if len(m) == 0 {
		return nil, errors.New("No Template Found")
	}
	var t *Template
	t, err = m2t(m[0])
	return t, err
}

func GetTemplateByTemplateName(name string) (*Template, error) {
	query := `{"selector" : {"template_name" : "` + name + `" }}`
	m, err := template_DB.QueryJSON(query)
	if err != nil {
		return nil, err
	}
	if len(m) == 0 {
		return nil, errors.New("No Template Found")
	}
	var t *Template
	t, err = m2t(m[0])
	return t, err
}

func CreateTemplate() (*Template, error) {
	id := utils.GenerateID()
	t := Template{
		_id:           "",
		_rev:          "",
		template_id:   id,
		template_name: "New Template",

		elements: []*Element{},
	}
	StoreTemplate(&t)
	return &t, nil
}

func (t *Template) AddElement() (*Element, error) {
	e := Element{
		element_id:     utils.GenerateID(),
		content:        "room_name",
		content_static: false,
		x:              0,
		y:              0,
		w:              0,
		h:              0,
		color:          "white",
		fill_color:     "black",
		font_size:      16,
		pixel_size:     1,
		style:          "bold",
		form:           "text",
	}
	fmt.Println("CREATED A NEW ELEMENT", e)
	t.elements = append(t.elements, &e)
	t.Save()
	return &e, nil
}

func (t *Template) Save() error {
	m := t2m(t)
	_, rev, err := template_DB.Save(m, nil)
	t._rev = rev
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func StoreTemplate(t *Template) error {
	m := t2m(t)
	delete(m, query.DB_ID)
	delete(m, query.DB_REV)
	id, rev, err := template_DB.Save(m, nil)
	if err != nil {
		log.Fatal(err)
		return err
	}

	t._id = id
	t._rev = rev
	return nil
}

func (t *Template) Delete() error {
	err := template_DB.Delete(t._id)
	return err
}

func (t *Template) GetElements() *[]*Element {
	return &t.elements
}

func (t *Template) GetName() string {
	return t.template_name
}

func (t *Template) GetID() string {
	return t.template_id
}

/**
Getter/Setter Element
*/

func (e *Element) GetElementID() string {
	return e.element_id
}

func (e *Element) GetContent() string {
	return e.content
}

func (e *Element) GetContentStatic() bool {
	return e.content_static
}

func (e *Element) GetX() int {
	return e.x
}

func (e *Element) GetY() int {
	return e.y
}

func (e *Element) GetW() int {
	return e.w
}

func (e *Element) GetH() int {
	return e.h
}

func (e *Element) GetColor() string {
	return e.color
}

func (e *Element) GetFillColor() string {
	return e.fill_color
}

func (e *Element) GetFontSize() int {
	return e.font_size
}

func (e *Element) GetPixelSize() int {
	return e.pixel_size
}

func (e *Element) GetStyle() string {
	return e.style
}

func (e *Element) GetForm() string {
	return e.form
}

func Check_Template(id string) (bool, *Template, error) {
	t, err := GetTemplateByTemplateID(id)
	if err != nil {
		return false, nil, err
	}
	return true, t, nil
}

func GetAllTemplates() ([]*Template, error) {
	ids, err := template_DB.DocIDs()
	if err != nil {
		return nil, err
	}
	t := []*Template{}
	for _, i := range ids {
		m, _ := template_DB.Get(i, nil)

		tmplt, _ := m2t(m)

		t = append(t, tmplt)

	}
	return t, err
}

func (t *Template) Remove() {
	template_DB.Delete(t._id)
}

func (t *Template) CheckElement(eid string) (bool, int) {
	if len(eid) == 0 {
		return false, -1
	}

	for i, k := range t.elements {
		fmt.Println(eid)
		if eid == k.element_id {
			fmt.Println(k.element_id)
			return true, i
		}
	}
	return false, -1
}

func (t *Template) RemoveElementByIndex(i int) error {
	if i > len(t.elements) {
		return errors.New("Element index out of Bounce")
	}
	t.elements = append(t.elements[:i], t.elements[i+1:]...)
	t.Save()
	return nil
}

func (t *Template) SetName(name string) {
	t.template_name = name
	t.Save()
}

func (t *Template) SaveSettings(
	index int,
	tname string,
	eid string,
	econtent string,
	estatic bool,
	x int,
	y int,
	w int,
	h int,
	color string,
	fill_color string,
	font_size int,
	pixel_size int,
	style string,
	form string) {
	t.template_name = tname
	t.elements[index].content = econtent
	t.elements[index].content_static = estatic
	t.elements[index].color = color
	t.elements[index].fill_color = fill_color
	t.elements[index].x = x
	t.elements[index].y = y
	t.elements[index].w = w
	t.elements[index].h = h
	t.elements[index].font_size = font_size
	t.elements[index].pixel_size = pixel_size
	t.elements[index].style = style
	t.elements[index].form = form
	fmt.Println(t)
	fmt.Println(t.elements[0])
	t.Save()
}
