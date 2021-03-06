package ibrest

import "errors"

//ErrNotRunning ib-resapi is not runnning
var ErrNotRunning error = errors.New("ib-restapi is not running make sure to call Start() first")

//ErrInvalidIP Invlid IPv4
var ErrInvalidIP error = errors.New("IP is not a valid IPv4")

//ErrInvalidPort Invlid Port Number
var ErrInvalidPort error = errors.New("Port is not a valid port number")

//ErrEndpointNotSet Endpoint has not been set
var ErrEndpointNotSet error = errors.New("Endpoint not set make sure to SetEndpoint() first")

//ErrCantConnect Connection to cpw couldnt be made
var ErrCantConnect error = errors.New("Can't connect to cpw")

//ErrCantAuthenticate Lost connection to cpw
var ErrCantAuthenticate error = errors.New("Unable to authenticate session. PLease try restarting your cpw")

//ErrUnkownResponseCode Unkown response code from server
var ErrUnkownResponseCode error = errors.New("Unkown response code from server")

//ErrInitSession cpw not initiated
var ErrInitSession error = errors.New("Please initiate the cpw session first")

//ErrNotAuthenticated cpw not authenticated
var ErrNotAuthenticated error = errors.New("The cpw is not currently authenticated, make sure you are logged into cpw")
