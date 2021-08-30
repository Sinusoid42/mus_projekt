package controller

/**

The package controller


This package is the main Package for answering HTTP Requests



This package divides multiple different controller paths into different readable files, such
that the readablilty is eased up

SubControllers:


	api.go

		Handles the multiplexing of api calls, when working with rooms, templates, arduinos or the menu
		Requires corresponding admin access by the user, which is trying to access the content provided on this page
		Arduinos have themselves seleceted paths reserved such that the arduinos can be dynamicly attached to the service which is provided by the server


	arduino.go

	@author ben
*/
