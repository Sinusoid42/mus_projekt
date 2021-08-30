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

 This js file contains default protocol defintions
 defining website interactions with the server

 querydefinitions, and http method extentions are defined
 with this js file in order to accommodate for different protocol
 needs when the webfrontend is communicating with the server

 see different definitions and comments for further notice

 @author benwinter (09.07.2021)


 */


/**
 * The http const is providing more
 * useful restful methods to the standard http protocol
 * to extend the functionality of the default header protocol
 *
 * since the header of http is string (utf-8) based, a extension
 * of the provided methods GPPD (get, post, put, delete) eg. CRUD within
 * the database the additional methods defined by method_extensions {me}
 * can be matched when communicating with the server
 *
 * @type {{me: {}}}
 */


const http = {

    /**
     * Extends the default http methods
     *
     * me : method extensions
     */

    me : {

        /**
         * http.CHECK
         *
         * Method extension to return
         * a check for the desired data
         *
         * server responds with json format
         * containing information of the
         * existance of the data or availabillty
         */
        check : "CHECK",

        /**
         * The regular http post method
         *
         * usability can be read in the restful dissertation
         * by roy fielding
         *
         * is used when data is send to the server in order
         * to create a new dataset
         */
        post : "POST",


        /**
         * The regular http get method
         *
         * has different use cases defined by the api
         * for each site
         */
        get : "GET",


        select : "SELECT",


        delete : "DELETE",

        put : "PUT",

        menu : "MENU",

        /**
         * The Fetch method extention
         *
         */
        fetch : "FETCH",



    },
    /**
     * ?<queries>:<values>
     *
     * for the standard http protocol, queries define
     * possible request parameters the server can answer to
     * and reacts for
     *
     *
     */
    q : {
        user_email : "user_email_address",
        user_exist : "user_exist",
        user_name : "user_name",
        user_password : "user_password",
        success : "success",
        authenticated : "auth",
        access : "admin_level",

        room_id : "room_id",
        room_name : "room_name",
        room_name_misc : "room_name_misc",
        room_bookable :"room_bookable",
        room_max_capacity : "room_max_capacity",
        room_max_duration : "room_max_booking_duration",
        room_location : "room_location",
        room_arduino_available : "room_arduino_available",

        location_room_floor_level : "location_room_floor_level",
        location_room_corridor : "location_room_corridor",
        location_room_number : "location_room_number",
        location_room_name : "location_room_name",


        template_id : "template_id",
        template_name : "template_name",
        template_elements : "elements",
        template_element_id : "element_id",
        template_element_content : "content",
        template_element_content_static : "content_static",
        template_element_position_x : "x",
        template_element_position_y : "y",
        template_element_position_w : "w",
        template_element_position_h : "h",
        template_element_color : "color",
        template_element_fill_color : "fill_color",
        template_element_font_size : "font_size",
        template_element_pixel_size : "pixel_size",
        template_element_style : "pixel_style",
        template_element_form : "form",


        booking_id : "booking_id",
        booking_user_id : "booking_user_id",
        booking_room_id : "booking_room_id",
        booking_occupancy : "booking_occupancy",
        booking_topic : "booking_topic",
        booking_description : "booking_description",

        booking_time_year : "booking_time_year",
        booking_time_week : "booking_time_week",
        booking_time_day : "booking_time_day",
        booking_time_month : "booking_time_month",

        booking_time_start : "booking_time_start",
        booking_time_end : "booking_time_end",

        booking_time_reoccuring : "booking_time_reoccuring",
        booking_time_end_year : "booking_time_end_year",
        booking_time_end_week : "booking_time_end_week",
        booking_time_end_month : "booking_time_end_month",
        booking_time_end_day : "booking_time_end_day",

        calender_year : "selected_calender_year",
        calender_week : "selected_calender_week",
        calender_day : "selected_calender_day",
        calender_hour : "selected_calender_hour",
    },
    /*
        path defintions
     */
    p : {
        login : "/login",
        logout : "/logout",
        register : "/register",
        arduino_api : "/api/arduinos",
        password_forget : "/pw_reset",
        confirmation : "/confirmation",
        room_api : "/api/rooms",
        template_api : "/api/templates",
        upgrade_api : "/api/upgrades",
        users : "/users",
        bookings : "/bookings",
        rooms : "/rooms",
    },
    a : {
        address : "",
        port : "",
    },


    putServerData(server_data_json){
        console.log(server_data_json)
        http.a.address = server_data_json[api.SERVER_ADDRESS_KEY]
        http.a.port = server_data_json[api.SERVER_PORT_KEY]
    }

}

/**
 * Helper Methods
 * @param e
 * @returns {HTMLElement}
 */
const $ = (e) => {return document.getElementById(e)}

const $$ = (e) => {return document.createElement(e)}


/**
 * The api const is providing communication
 * definitions when communiating with the server
 * and exchanging data
 *
 * to ease the use of all communication with the server
 * variable definitions are put here, so they are easily
 * accessible and are standardized, easily editable and
 * scaleable throughout the entire project
 *
 *
 * @type {{resolveServerAddress: Promise<unknown>, SERVER_ADDRESS_KEY: string, SERVER_PORT_KEY: string}}
 */

const api = {

    SERVER_ADDRESS_KEY : "server_address",
    SERVER_PORT_KEY : "server_port",


    resolveServerAddress : new Promise((resolve) => {
        const request = new XMLHttpRequest();
        request.open("GET", "/api/server", true);
        request.send();
        request.onreadystatechange = () => {
            if (request.readyState == XMLHttpRequest.DONE){
                resolve(JSON.parse(request.response));
            }
        }
    }),
}



/** - - - - - - - - - - - - - - - - - - - - -

 ColorSceme definitions by David Martin Karg

 @author benwinter

 - - - - - - - - - - - - - - - - - - - - -

 lb : light-blue
 sb : secondary-blue
 pb : primary-blue
 db : dark-blue

 lr : light-red
 sr : secondary-red
 pr : primary-red
 dr : dark-red
 */

const color_sceme = {

    lb : "#CEE3F2",
    sb : "#A7BFD1",
    pb : "#4C657C",
    db : "#273A4C",

    lr : "#FFCCD2",
    sr : "#ED8A98",
    pr : "#DB4F60",
    dr : "#C1303F",
}






























