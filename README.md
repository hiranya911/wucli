# wucli
Command-line Interface for Weather Underground API (http://www.wunderground.com/)

NAME
	wucli -- Weather Underground Command-line Interface

USAGE
	wucli [location/keyword]

DESCRIPTION
	Run wucli without any arguments to check the weather of your current location (determined
	automatically from your IP address).

	Optionally you can run wucli with a location argument to get weather data related to a specific
	city or area. Following types of location arguments are supported as of now:

	  * US state/city [e.g. CA/San Francisco]
	  * US zipcode [e.g. 93106]
	  * country/city [e.g. Sri Lanka/Colombo]
	  * latitude,longitude [e.g. 37.8,-122.4]
	  * airport code [e.g. KJFK]

	You can also enter a simple keyword or a phrase as a search term [e.g. Belgrade]. In this case
	wucli will return a list of matching locations with their unique zmw codes. You may rerun wucli
	with a zmw code as a location argument, to indicate which location that you are interested in.

	If the location argument or the keyword contains spaces, enclose the argument within quotes.
	Alternatively, you can replace the spaces with underscores (_).

BUILD
	Execute the build.sh script. This will build the code, and copy the resulting binary to /usr/bin
	directory (tested on OS-X Mountain Lion with Golang 1.2.1).