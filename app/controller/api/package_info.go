package api

/**
The controller/api package

The main workload abstraction this package handles is the multiplexing
of all incoming API calls

Funneling all internal Server API Paths through a single Mux enalbles
easier capsuling and code reusage

The api itself is consisting of Admin-Priviledge (or similar) access
to the API Service, the ROS Group Web App has to offer

The API consists of /api/rooms, where the service administrator
can handle the addition and removal of office rooms, select
already existing rooms and alter their settings according to different
settings

//Todos are resoluted
//TODO:
	API Rooms differentiate between user priviledges when handling the input
	Enable Service Admins, Professors etc. to create Bookings for rooms,
	even if those are publicly not available

	Create USER Page
			=> show all bookings of the user
			=> make bookings clickable, when clicked, show settings for the booking corresponding to  the selected one

				=> hyperlink to the room
				=> hyperlink to the user, when a booking within the settings was clicked
			=>Integrate "Upgrade Account" Button, to create a new Ticket for upgrading a account



	Create API Upgrade Account Page
		=> list all tickets of possible upgradeable accounts

		=> buttons [reject, accept] + value for access type


	Adjust dynamic data for the arduino     => content has to be dynamic



	@author ben

*/
