package gala

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

const IconifyVersion = ""

var (
	hAttr   = g.Attr("height", "1em")
	viewbox = g.Attr("viewbox", "0 0 256 256")
)

func Add(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaAdd0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" stroke-width="16"><circle id="galaAdd1" cx="128" cy="128" r="112"/><path id="galaAdd2" d="M 79.999992,128 H 176.0001"/><path id="galaAdd3" d="m 128.00004,79.99995 v 96.0001"/></g>`), g.Group(children),
	)
}

func Airplay(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaAirplay0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" stroke-width="16"><path id="galaAirplay1" d="M 64,192 H 47.999992 m 0,0 c -17.728,0 -32,-14.272 -32,-32 V 47.999993 c 0,-17.728 14.272,-32 32,-32 H 208.00001 c 17.728,0 32,14.272 32,32 V 160 c 0,17.728 -14.272,32 -32,32"/><path id="galaAirplay2" d="m 64,240 64,-80 64,80 z"/><path id="galaAirplay3" d="M 208,192 H 192"/></g>`), g.Group(children),
	)
}

func Apple(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaApple0"><g id="galaApple1" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" stroke-width="4.244" transform="translate(.331 .207) scale(3.76975)"><path id="galaApple2" d="m 33.866666,21.166666 c -4.233333,1e-6 -25.465872,-8.433732 -25.465872,12.732936"/><path id="galaApple3" d="m 8.400794,33.899602 c 0,21.166665 12.743914,29.710184 16.977248,29.710184 4.233334,0 4.255291,-4.244312 8.488624,-4.244312"/><path id="galaApple4" d="m 33.866665,21.166666 c 4.233334,1e-6 25.465874,-8.433732 25.465874,12.732936"/><path id="galaApple5" d="m 59.354497,33.866666 c 0,21.166666 -12.765873,29.74312 -16.999207,29.74312 -4.233334,0 -4.25529,-4.244312 -8.488624,-4.244312"/><path id="galaApple6" d="m 33.866666,21.166666 c 0,-12.6999995 -8.488624,-16.9772483 -8.488624,-16.9772483"/><path id="galaApple7" d="m 33.866666,21.166666 c 0,-16.9333325 8.499603,-12.7439153 12.732936,-16.9772483"/><path id="galaApple8" d="m 33.866666,21.166666 c 12.7,-4.233333 16.966269,-8.510583 12.732936,-16.9772483"/></g></g>`), g.Group(children),
	)
}

func Bag(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaBag0" fill="none" stroke-width="16"><path id="galaBag1" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" d="m 64,80 h 128 c 16,0 29.33333,16 32,32 l 16,96 c 2.66807,16.00842 -16,32 -32,32 H 48 C 32,240 13.33193,224.00842 16,208 L 32,112 C 34.666667,96 48,80 64,80 Z"/><path id="galaBag2" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" d="M 80,112 V 63.814079"/><path id="galaBag3" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" d="m 176,64 v 48"/><path id="galaBag4" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" d="M 19.090159,191.31828 H 236.90984"/><path id="galaBag5" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" d="M 176,64 C 176,48 166.70076,30.94703 151.90703,22.405869 137.1133,13.86471 118.88668,13.86471 104.09296,22.40587 89.299233,30.947031 80.000002,48 80,64"/><rect id="galaBag6" width="80" height="16" x="16" y="240" stroke="none"/></g>`), g.Group(children),
	)
}

func Bell(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaBell0"><g id="galaBell1" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-width="4.233" transform="translate(16) scale(3.77953)"><path id="galaBell2" stroke-opacity="1" d="M 33.866678,59.266666 A 4.2333331,4.2333331 0 0 1 29.633345,63.5"/><path id="galaBell3" stroke-opacity="1" d="M -25.4,59.266666 A 4.2333331,4.2333331 0 0 1 -29.633333,63.5" transform="scale(-1 1)"/><path id="galaBell4" stroke-opacity="1" d="m 25.400031,55.033482 v 4.233333"/><path id="galaBell5" stroke-opacity="1" d="m 33.866698,55.033482 -1.9e-5,4.233332"/><path id="galaBell6" d="m 55.033333,50.8 c -8.466667,4.233333 -21.166667,4.233333 -25.4,4.233333"/><path id="galaBell7" stroke-opacity="1" d="m 55.033333,50.8 c 0,-8.466667 -8.466667,-4.233336 -8.466667,-25.400002 0,-9.466021 -8.466666,-16.9330112 -12.699988,-16.9330112"/><path id="galaBell8" d="M 4.233334,50.8 C 12.7,55.033333 25.4,55.033333 29.633333,55.033333"/><path id="galaBell9" stroke-opacity="1" d="m 4.2333451,50.799681 c 0,-8.466669 8.4666669,-4.233338 8.4666669,-25.400003 0,-9.46602 8.466666,-16.9330108 12.699988,-16.9330108"/><path id="galaBella" stroke-opacity="1" d="m -25.400012,-8.4666872 a 4.2333331,4.2333331 0 0 1 -4.233333,4.2333331" transform="scale(-1)"/><path id="galaBellb" stroke-opacity="1" d="m 33.866666,-8.4666672 a 4.2333331,4.2333331 0 0 1 -4.233333,4.2333331" transform="scale(1 -1)"/></g></g>`), g.Group(children),
	)
}

func Book(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaBook0"><g id="galaBook1" stroke-dasharray="none" stroke-miterlimit="4" stroke-width="4.234" transform="translate(.005 .01) scale(3.77938)"><g id="galaBook2" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-width="4.234" transform="translate(4.233)"><path id="galaBook3" stroke-opacity="1" d="M 55.033333,63.5 12.7,63.499999"/><path id="galaBook4" d="M 4.2333336,-55.033333 A 8.4666662,8.4666662 0 0 1 12.7,-63.499999" transform="scale(1 -1)"/><path id="galaBook5" d="M 4.2333336,55.033333 A 8.4666662,8.4666662 0 0 1 12.7,46.566667"/><path id="galaBook6" stroke-opacity="1" d="m 55.033333,46.566667 -42.333332,-1e-6"/><path id="galaBook7" d="M 4.2333336,12.7 A 8.4666662,8.4666662 0 0 1 12.7,4.2333336"/><path id="galaBook8" stroke-opacity="1" d="m 55.033333,4.233334 -42.333332,-1e-6"/><path id="galaBook9" stroke-opacity="1" d="m 4.2333334,12.7 2e-7,42.333333"/><path id="galaBooka" stroke-opacity="1" d="M 12.7,46.566666 V 4.230783"/><path id="galaBookb" stroke-opacity="1" d="M 55.033333,4.2333348 V 46.566667"/><path id="galaBookc" stroke-opacity="1" d="m 55.033333,46.566667 c -2.662642,5.559118 -2.809222,11.198865 0,16.933333"/></g></g></g>`), g.Group(children),
	)
}

func Brochure(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaBrochure0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" stroke-width="16"><path id="galaBrochure1" d="M 16.110084,16.110084 H 160.04176 L 96.072129,64.087313"/><path id="galaBrochure2" d="m 16.110084,16.110084 -2e-6,175.916496 h 79.962047"/><path id="galaBrochure3" d="M 96.072132,64.087313 H 240.00381 V 240.0038 H 96.072129 l 3e-6,-175.916487"/><path id="galaBrochure4" d="M 160.04176,16.110084 V 64.087313"/></g>`), g.Group(children),
	)
}

func Calendar(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaCalendar0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-miterlimit="4" stroke-width="16"><path id="galaCalendar1" stroke-linecap="butt" stroke-linejoin="miter" stroke-opacity="1" d="M 31.999978,31.999961 H 224.00004"/><path id="galaCalendar2" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="m 15.999975,47.999965 -3e-6,176.000055"/><path id="galaCalendar3" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="M 240.00002,47.999965 V 224.00002"/><path id="galaCalendar4" stroke-linecap="round" stroke-linejoin="round" d="m 31.999978,31.999961 c -8.836576,0 -16.000003,7.163446 -16.000003,16.000004"/><path id="galaCalendar5" stroke-linecap="round" stroke-linejoin="round" d="m 224.00004,31.999961 c 8.83657,-4e-6 15.99998,7.163443 15.99998,16.000004"/><path id="galaCalendar6" stroke-linecap="butt" stroke-linejoin="miter" stroke-opacity="1" d="M 224.00004,240.00002 H 31.999978"/><path id="galaCalendar7" stroke-linecap="round" stroke-linejoin="round" d="m 224.00004,240.00002 a 16.000004,16.000004 0 0 0 15.99998,-16"/><path id="galaCalendar8" stroke-linecap="round" stroke-linejoin="round" d="m 31.999978,240.00002 a 16.000004,16.000004 0 0 1 -16.000011,-16"/><path id="galaCalendar9" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="M 128.00001,47.999965 V 15.999962"/><path id="galaCalendara" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="M 160.00003,47.999965 V 15.999962"/><path id="galaCalendarb" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="M 192.00002,47.999965 V 15.999962"/><path id="galaCalendarc" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="M 95.999985,47.999965 V 15.999962"/><path id="galaCalendard" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="M 64.000001,47.999965 V 15.999962"/><path id="galaCalendare" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="M 15.999975,95.999972 H 240.00002"/></g>`), g.Group(children),
	)
}

func Chart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaChart0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" stroke-width="16"><rect id="galaChart1" width="224" height="224" x="16" y="16" ry="32"/><path id="galaChart2" d="M 160.00003,192.00003 V 111.99998"/><path id="galaChart3" d="M 192.00002,192.00003 V 63.999974"/><path id="galaChart4" d="m 63.999979,192.00003 v -32"/><path id="galaChart5" d="m 95.99997,128 v 64.00003"/><path id="galaChart6" d="m 128,144 v 48.00003"/></g>`), g.Group(children),
	)
}

func Chat(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaChat0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-miterlimit="4" stroke-opacity="1" stroke-width="16"><path id="galaChat1" stroke-linejoin="miter" d="M 48.004588,16.00078 H 207.99487"/><path id="galaChat2" stroke-linejoin="round" d="M 240.00422,-48.000717 A 31.999937,31.999937 0 0 1 208.00429,-16.00078" transform="scale(1 -1)"/><path id="galaChat3" stroke-linejoin="round" d="M -16.004652,-48.000717 A 31.999937,31.999937 0 0 1 -48.004588,-16.00078" transform="scale(-1)"/><path id="galaChat4" stroke-linejoin="round" d="m -16.004652,160.0005 a 31.999937,31.999937 0 0 1 -31.999936,31.99994" transform="scale(-1 1)"/><path id="galaChat5" stroke-linejoin="round" d="m 240.00422,160.0005 a 31.999937,31.999937 0 0 1 -31.99993,31.99994"/><path id="galaChat6" stroke-linejoin="round" d="M 16.004649,48.00072 V 160.00051"/><path id="galaChat7" stroke-linejoin="round" d="M 240.00422,48.00072 V 160.00051"/><path id="galaChat8" stroke-linejoin="round" d="m 48.004588,192.00044 h 31.999937 l 47.999905,47.99991 47.99991,-47.99991 h 31.99994"/></g>`), g.Group(children),
	)
}

func Clock(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaClock0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-width="16"><path id="galaClock1" stroke-opacity="1" d="M 80.039519,211.06969 64.056987,239.90745"/><path id="galaClock2" stroke-opacity="1" d="m 175.96058,211.06969 15.98254,28.83776"/><circle id="galaClock3" cx="128" cy="128.007" r="95.915"/><path id="galaClock4" d="M 35.294352,102.43866 A 47.957299,47.957299 0 0 1 17.212686,53.792007 47.957299,47.957299 0 0 1 53.990946,17.175027 47.957299,47.957299 0 0 1 102.55767,35.470309"/><path id="galaClock5" stroke-opacity="1" d="m 127.99967,32.092482 3.8e-4,-15.985761"/><path id="galaClock6" stroke-opacity="1" d="M 128.00005,80.049788 V 128.00708"/><path id="galaClock7" stroke-opacity="1" d="m 128.00005,128.00708 33.91093,33.91093"/><path id="galaClock8" d="M 220.70575,102.43866 A 47.957299,47.957299 0 0 0 238.78742,53.792007 47.957299,47.957299 0 0 0 202.00916,17.175027 47.957299,47.957299 0 0 0 153.44244,35.470309"/></g>`), g.Group(children),
	)
}

func Copy(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaCopy0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-width="16"><rect id="galaCopy1" width="144" height="144" x="16" y="16" ry="32"/><path id="galaCopy2" d="M 95.999712,127.9996 A 31.999891,31.999891 0 0 1 127.9996,95.999712"/><path id="galaCopy3" d="m -239.99922,127.9996 a 31.999891,31.999891 0 0 1 31.99989,-31.999888" transform="scale(-1 1)"/><path id="galaCopy4" d="m -239.99922,-207.99933 a 31.999891,31.999891 0 0 1 31.99989,-31.99989" transform="scale(-1)"/><path id="galaCopy5" d="M 95.999712,-207.99933 A 31.999891,31.999891 0 0 1 127.9996,-239.99922" transform="scale(1 -1)"/><path id="galaCopy6" stroke-opacity="1" d="m 159.99949,239.99923 h 15.99995"/><path id="galaCopy7" stroke-opacity="1" d="m 159.99949,95.9997 h 15.99995"/><path id="galaCopy8" stroke-opacity="1" d="m 95.999709,159.99949 v 15.99995"/><path id="galaCopy9" stroke-opacity="1" d="m 239.99923,159.99949 v 15.99995"/></g>`), g.Group(children),
	)
}

func Data(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaData0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-width="16"><path id="galaData1" d="M 239.98507,55.993592 A 111.98507,39.994664 0 0 1 128,95.988256 111.98507,39.994664 0 0 1 16.01493,55.993592 111.98507,39.994664 0 0 1 128,15.998927 111.98507,39.994664 0 0 1 239.98507,55.993592 Z"/><path id="galaData2" d="m 239.98507,199.97441 a 111.98507,39.994664 0 0 1 -55.99253,34.63639 111.98507,39.994664 0 0 1 -111.985079,0 111.98507,39.994664 0 0 1 -55.992531,-34.6364"/><path id="galaData3" d="m 239.98507,151.9808 a 111.98507,39.994664 0 0 1 -55.99253,34.6364 111.98507,39.994664 0 0 1 -111.985079,-1e-5 A 111.98507,39.994664 0 0 1 16.01493,151.9808"/><path id="galaData4" d="m 239.98507,103.9872 a 111.98507,39.994664 0 0 1 -55.99253,34.6364 111.98507,39.994664 0 0 1 -111.985079,0 111.98507,39.994664 0 0 1 -55.992531,-34.6364"/><path id="galaData5" stroke-opacity="1" d="M 16.01493,55.99377 V 199.97441"/><path id="galaData6" stroke-opacity="1" d="M 239.98507,55.993592 V 199.97441"/></g>`), g.Group(children),
	)
}

func Display(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaDisplay0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-miterlimit="4" stroke-width="16"><path id="galaDisplay1" stroke-linecap="butt" stroke-linejoin="miter" stroke-opacity="1" d="M 32.045863,15.996967 H 223.95413"/><path id="galaDisplay2" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="m 16.053507,31.994108 2e-6,143.926412"/><path id="galaDisplay3" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="M 239.94649,31.989325 V 175.92052"/><path id="galaDisplay4" stroke-linecap="round" stroke-linejoin="round" d="M 32.045861,15.996971 A 15.992355,15.992355 0 0 0 16.053507,31.989325"/><path id="galaDisplay5" stroke-linecap="round" stroke-linejoin="round" d="m 223.95413,15.996971 a 15.992355,15.992355 0 0 1 15.99236,15.992354"/><path id="galaDisplay6" stroke-linecap="butt" stroke-linejoin="miter" stroke-opacity="1" d="M 223.95413,191.91287 H 32.04587"/><path id="galaDisplay7" stroke-linecap="round" stroke-linejoin="round" d="m 223.95413,191.91287 a 15.992354,15.992354 0 0 0 15.99236,-15.99235"/><path id="galaDisplay8" stroke-linecap="round" stroke-linejoin="round" d="M 32.04587,191.91287 A 15.992354,15.992354 0 0 1 16.053504,175.92052"/><path id="galaDisplay9" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="M 191.96942,239.88994 H 64.030574"/><path id="galaDisplaya" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="m 143.99235,191.91288 v 47.97706"/><path id="galaDisplayb" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="m 112.00764,191.91288 v 47.97706"/><path id="galaDisplayc" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="M 16.053512,159.92816 H 239.94649"/></g>`), g.Group(children),
	)
}

func Editor(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaEditor0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" stroke-width="16"><path id="galaEditor1" d="m 16,64 224.93778,0.09256"/><path id="galaEditor2" d="M 48,16 H 207.91114 C 225.62929,16 240,30.281849 240,48 v 160 c 0,17.71816 -14.28185,32 -32,32 H 48 C 30.281848,240 16.069099,225.73073 16.06221,208.01257 L 16,48 C 15.993112,30.281851 30.281848,16 48,16 Z"/><path id="galaEditor3" d="M 191.96444,64.092555 192,16"/><path id="galaEditor4" d="M 48.044437,112.06589 H 80.02666"/><path id="galaEditor5" d="M 48.044437,144.04812 H 175.97333"/><path id="galaEditor6" d="M 48.044437,176.03034 H 127.99999"/><path id="galaEditor7" d="M 48.044437,208.01256 H 80.02666"/></g>`), g.Group(children),
	)
}

func File(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaFile0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-miterlimit="4" stroke-opacity="1" stroke-width="15.992"><path id="galaFile1" stroke-linejoin="miter" d="M 32,48 V 207.9236"/><path id="galaFile2" stroke-linejoin="round" d="M 224,96 V 208"/><path id="galaFile3" stroke-linejoin="round" d="m 64,16 h 80"/><path id="galaFile4" stroke-linejoin="miter" d="M 64,240 H 192"/><path id="galaFile5" stroke-linejoin="round" d="m 224,208 c 0.0874,15.98169 -16,32 -32,32"/><path id="galaFile6" stroke-linejoin="round" d="m -32,208 c -10e-7,16 -16,32 -32,32" transform="scale(-1 1)"/><path id="galaFile7" stroke-linejoin="round" d="M -32,-47.976784 C -32,-32 -48,-16.356322 -63.999997,-16.000002" transform="scale(-1)"/><path id="galaFile8" stroke-linejoin="round" d="M 223.91257,96.071779 144,16"/><path id="galaFile9" stroke-linejoin="round" d="m -144,64 c -0.0492,15.912926 -16.06452,31.999995 -32,32" transform="scale(-1 1)"/><path id="galaFilea" stroke-linejoin="round" d="M 144,64 V 16"/><path id="galaFileb" stroke-linejoin="round" d="m 176,96 h 48"/></g>`), g.Group(children),
	)
}

func FileCodeOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaFileCode10" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-miterlimit="4" stroke-opacity="1"><path id="galaFileCode11" stroke-linejoin="miter" stroke-width="15.992" d="M 32,48 V 207.9236"/><path id="galaFileCode12" stroke-linejoin="round" stroke-width="15.992" d="M 224,96 V 208"/><path id="galaFileCode13" stroke-linejoin="round" stroke-width="15.992" d="m 64,16 h 80"/><path id="galaFileCode14" stroke-linejoin="miter" stroke-width="15.992" d="M 64,240 H 192"/><path id="galaFileCode15" stroke-linejoin="round" stroke-width="15.992" d="m 224,208 c 0.0874,15.98169 -16,32 -32,32"/><path id="galaFileCode16" stroke-linejoin="round" stroke-width="15.992" d="m -32,208 c -10e-7,16 -16,32 -32,32" transform="scale(-1 1)"/><path id="galaFileCode17" stroke-linejoin="round" stroke-width="15.992" d="M -32,-47.976784 C -32,-32 -48,-16.356322 -63.999997,-16.000002" transform="scale(-1)"/><path id="galaFileCode18" stroke-linejoin="round" stroke-width="15.992" d="M 223.91257,96.071779 144,16"/><path id="galaFileCode19" stroke-linejoin="round" stroke-width="15.992" d="m -144,64 c -0.0492,15.912926 -16.06452,31.999995 -32,32" transform="scale(-1 1)"/><path id="galaFileCode1a" stroke-linejoin="round" stroke-width="15.992" d="M 144,64 V 16"/><path id="galaFileCode1b" stroke-linejoin="round" stroke-width="15.992" d="m 176,96 h 48"/><path id="galaFileCode1c" stroke-linejoin="round" stroke-width="16" d="M 96,208 64,176 96,144"/><path id="galaFileCode1d" stroke-linejoin="round" stroke-width="16" d="m 128,208 32,-32 -32,-32"/></g>`), g.Group(children),
	)
}

func FileCodeTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaFileCode20" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-miterlimit="4" stroke-opacity="1"><path id="galaFileCode21" stroke-linejoin="miter" stroke-width="15.992" d="M 32,48 V 207.9236"/><path id="galaFileCode22" stroke-linejoin="round" stroke-width="15.992" d="M 224,96 V 208"/><path id="galaFileCode23" stroke-linejoin="round" stroke-width="15.992" d="m 64,16 h 80"/><path id="galaFileCode24" stroke-linejoin="miter" stroke-width="15.992" d="M 64,240 H 192"/><path id="galaFileCode25" stroke-linejoin="round" stroke-width="15.992" d="m 224,208 c 0.0874,15.98169 -16,32 -32,32"/><path id="galaFileCode26" stroke-linejoin="round" stroke-width="15.992" d="m -32,208 c -10e-7,16 -16,32 -32,32" transform="scale(-1 1)"/><path id="galaFileCode27" stroke-linejoin="round" stroke-width="15.992" d="M -32,-47.976784 C -32,-32 -48,-16.356322 -63.999997,-16.000002" transform="scale(-1)"/><path id="galaFileCode28" stroke-linejoin="round" stroke-width="15.992" d="M 223.91257,96.071779 144,16"/><path id="galaFileCode29" stroke-linejoin="round" stroke-width="15.992" d="m -144,64 c -0.0492,15.912926 -16.06452,31.999995 -32,32" transform="scale(-1 1)"/><path id="galaFileCode2a" stroke-linejoin="round" stroke-width="15.992" d="M 144,64 V 16"/><path id="galaFileCode2b" stroke-linejoin="round" stroke-width="15.992" d="m 176,96 h 48"/><path id="galaFileCode2c" stroke-linejoin="round" stroke-width="16" d="m 96,208 c -32,0 0,-32 -32,-32 32,0 0,-32 32,-32"/><path id="galaFileCode2d" stroke-linejoin="round" stroke-width="16" d="m 128,208 c 32,0 0,-32 32,-32 -32,0 0,-32 -32,-32"/></g>`), g.Group(children),
	)
}

func FileDocument(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaFileDocument0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-miterlimit="4" stroke-opacity="1"><path id="galaFileDocument1" stroke-linejoin="miter" stroke-width="15.992" d="M 32,48 V 207.9236"/><path id="galaFileDocument2" stroke-linejoin="round" stroke-width="15.992" d="M 224,96 V 208"/><path id="galaFileDocument3" stroke-linejoin="round" stroke-width="15.992" d="m 64,16 h 80"/><path id="galaFileDocument4" stroke-linejoin="miter" stroke-width="15.992" d="M 64,240 H 192"/><path id="galaFileDocument5" stroke-linejoin="round" stroke-width="15.992" d="m 224,208 c 0.0874,15.98169 -16,32 -32,32"/><path id="galaFileDocument6" stroke-linejoin="round" stroke-width="15.992" d="m -32,208 c -10e-7,16 -16,32 -32,32" transform="scale(-1 1)"/><path id="galaFileDocument7" stroke-linejoin="round" stroke-width="15.992" d="M -32,-47.976784 C -32,-32 -48,-16.356322 -63.999997,-16.000002" transform="scale(-1)"/><path id="galaFileDocument8" stroke-linejoin="round" stroke-width="15.992" d="M 223.91257,96.071779 144,16"/><path id="galaFileDocument9" stroke-linejoin="round" stroke-width="15.992" d="m -144,64 c -0.0492,15.912926 -16.06452,31.999995 -32,32" transform="scale(-1 1)"/><path id="galaFileDocumenta" stroke-linejoin="round" stroke-width="15.992" d="M 144,64 V 16"/><path id="galaFileDocumentb" stroke-linejoin="round" stroke-width="15.992" d="m 176,96 h 48"/><path id="galaFileDocumentc" stroke-linejoin="round" stroke-width="16" d="m 64,208 h 48"/><path id="galaFileDocumentd" stroke-linejoin="round" stroke-width="16" d="m 64,176 h 80"/><path id="galaFileDocumente" stroke-linejoin="round" stroke-width="16" d="m 64,144 h 48"/></g>`), g.Group(children),
	)
}

func FileError(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaFileError0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-miterlimit="4" stroke-opacity="1"><path id="galaFileError1" stroke-linejoin="miter" stroke-width="15.992" d="M 32,48 V 207.9236"/><path id="galaFileError2" stroke-linejoin="round" stroke-width="15.992" d="M 224,96 V 208"/><path id="galaFileError3" stroke-linejoin="round" stroke-width="15.992" d="m 64,16 h 80"/><path id="galaFileError4" stroke-linejoin="miter" stroke-width="15.992" d="M 64,240 H 192"/><path id="galaFileError5" stroke-linejoin="round" stroke-width="15.992" d="m 224,208 c 0.0874,15.98169 -16,32 -32,32"/><path id="galaFileError6" stroke-linejoin="round" stroke-width="15.992" d="m -32,208 c -10e-7,16 -16,32 -32,32" transform="scale(-1 1)"/><path id="galaFileError7" stroke-linejoin="round" stroke-width="15.992" d="M -32,-47.976784 C -32,-32 -48,-16.356322 -63.999997,-16.000002" transform="scale(-1)"/><path id="galaFileError8" stroke-linejoin="round" stroke-width="15.992" d="M 223.91257,96.071779 144,16"/><path id="galaFileError9" stroke-linejoin="round" stroke-width="15.992" d="m -144,64 c -0.0492,15.912926 -16.06452,31.999995 -32,32" transform="scale(-1 1)"/><path id="galaFileErrora" stroke-linejoin="round" stroke-width="15.992" d="M 144,64 V 16"/><path id="galaFileErrorb" stroke-linejoin="round" stroke-width="15.992" d="m 176,96 h 48"/><path id="galaFileErrorc" stroke-linejoin="round" stroke-width="16" d="m 64,208 48,-64"/><path id="galaFileErrord" stroke-linejoin="round" stroke-width="16" d="m 64,144 48,64"/></g>`), g.Group(children),
	)
}

func FileScript(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaFileScript0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-miterlimit="4" stroke-opacity="1"><path id="galaFileScript1" stroke-linejoin="miter" stroke-width="15.992" d="M 32,48 V 207.9236"/><path id="galaFileScript2" stroke-linejoin="round" stroke-width="15.992" d="M 224,96 V 208"/><path id="galaFileScript3" stroke-linejoin="round" stroke-width="15.992" d="m 64,16 h 80"/><path id="galaFileScript4" stroke-linejoin="miter" stroke-width="15.992" d="M 64,240 H 192"/><path id="galaFileScript5" stroke-linejoin="round" stroke-width="15.992" d="m 224,208 c 0.0874,15.98169 -16,32 -32,32"/><path id="galaFileScript6" stroke-linejoin="round" stroke-width="15.992" d="m -32,208 c -10e-7,16 -16,32 -32,32" transform="scale(-1 1)"/><path id="galaFileScript7" stroke-linejoin="round" stroke-width="15.992" d="M -32,-47.976784 C -32,-32 -48,-16.356322 -63.999997,-16.000002" transform="scale(-1)"/><path id="galaFileScript8" stroke-linejoin="round" stroke-width="15.992" d="M 223.91257,96.071779 144,16"/><path id="galaFileScript9" stroke-linejoin="round" stroke-width="15.992" d="m -144,64 c -0.0492,15.912926 -16.06452,31.999995 -32,32" transform="scale(-1 1)"/><path id="galaFileScripta" stroke-linejoin="round" stroke-width="15.992" d="M 144,64 V 16"/><path id="galaFileScriptb" stroke-linejoin="round" stroke-width="15.992" d="m 176,96 h 48"/><path id="galaFileScriptc" stroke-linejoin="round" stroke-width="16" d="M 64,208 96.000003,176 64,144"/><path id="galaFileScriptd" stroke-linejoin="round" stroke-width="16" d="m 112,208 h 32"/></g>`), g.Group(children),
	)
}

func FileSpreadsheet(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaFileSpreadsheet0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-miterlimit="4" stroke-opacity="1"><path id="galaFileSpreadsheet1" stroke-linejoin="miter" stroke-width="15.992" d="M 32,48 V 207.9236"/><path id="galaFileSpreadsheet2" stroke-linejoin="round" stroke-width="15.992" d="M 224,96 V 208"/><path id="galaFileSpreadsheet3" stroke-linejoin="round" stroke-width="15.992" d="m 64,16 h 80"/><path id="galaFileSpreadsheet4" stroke-linejoin="miter" stroke-width="15.992" d="M 64,240 H 192"/><path id="galaFileSpreadsheet5" stroke-linejoin="round" stroke-width="15.992" d="m 224,208 c 0.0874,15.98169 -16,32 -32,32"/><path id="galaFileSpreadsheet6" stroke-linejoin="round" stroke-width="15.992" d="m -32,208 c -10e-7,16 -16,32 -32,32" transform="scale(-1 1)"/><path id="galaFileSpreadsheet7" stroke-linejoin="round" stroke-width="15.992" d="M -32,-47.976784 C -32,-32 -48,-16.356322 -63.999997,-16.000002" transform="scale(-1)"/><path id="galaFileSpreadsheet8" stroke-linejoin="round" stroke-width="15.992" d="M 223.91257,96.071779 144,16"/><path id="galaFileSpreadsheet9" stroke-linejoin="round" stroke-width="15.992" d="m -144,64 c -0.0492,15.912926 -16.06452,31.999995 -32,32" transform="scale(-1 1)"/><path id="galaFileSpreadsheeta" stroke-linejoin="round" stroke-width="15.992" d="M 144,64 V 16"/><path id="galaFileSpreadsheetb" stroke-linejoin="round" stroke-width="15.992" d="m 176,96 h 48"/><path id="galaFileSpreadsheetc" stroke-linejoin="round" stroke-width="16" d="m 64,192 h 96"/><path id="galaFileSpreadsheetd" stroke-linejoin="round" stroke-width="16" d="m 64,160 h 96"/><path id="galaFileSpreadsheete" stroke-linejoin="round" stroke-width="16" d="m 80,144 v 64"/><path id="galaFileSpreadsheetf" stroke-linejoin="round" stroke-width="16" d="m 112,144 v 64"/><path id="galaFileSpreadsheetg" stroke-linejoin="round" stroke-width="16" d="m 144,144 v 64"/></g>`), g.Group(children),
	)
}

func FileText(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaFileText0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" stroke-width="16"><path id="galaFileText1" d="M 32,48 V 207.9236"/><path id="galaFileText2" d="M 224,96 V 208"/><path id="galaFileText3" d="m 64,16 h 80"/><path id="galaFileText4" d="M 64,240 H 192"/><path id="galaFileText5" d="m 224,208 c 0.0874,15.98169 -16,32 -32,32"/><path id="galaFileText6" d="m -32,208 c -10e-7,16 -16,32 -32,32" transform="scale(-1 1)"/><path id="galaFileText7" d="M -32,-47.976784 C -32,-32 -48,-16.356322 -63.999997,-16.000002" transform="scale(-1)"/><path id="galaFileText8" d="M 223.91257,96.071779 144,16"/><path id="galaFileText9" d="m -144,64 c -0.0492,15.912926 -16.06452,31.999995 -32,32" transform="scale(-1 1)"/><path id="galaFileTexta" d="M 144,64 V 16"/><path id="galaFileTextb" d="m 176,96 h 48"/><path id="galaFileTextc" d="m 80,208 h 32"/><path id="galaFileTextd" d="m 96,144 v 64"/><path id="galaFileTexte" d="m 64,144 h 64"/><path id="galaFileTextf" d="m 128,144 v 16"/><path id="galaFileTextg" d="m 64,144 v 16"/></g>`), g.Group(children),
	)
}

func Folder(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaFolder0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-miterlimit="4" stroke-width="16"><path id="galaFolder1" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="M 240.04121,64.000608 240,208 c -0.005,15.99988 -16.04134,32 -32,32 H 48 C 31.958138,240 15.996182,223.99988 16,208 l 0.04199,-175.999258"/><path id="galaFolder2" stroke-linecap="round" stroke-linejoin="round" d="m 224,48 c 8.83652,8e-6 15.99999,7.163485 16,16"/><path id="galaFolder3" stroke-linecap="butt" stroke-linejoin="miter" stroke-opacity="1" d="m 144.04154,48.000675 h 79.99974"/><path id="galaFolder4" stroke-linecap="round" stroke-linejoin="round" d="m 106.38483,16.000772 a 15.999945,15.999945 0 0 1 13.85635,7.999985"/><path id="galaFolder5" stroke-linecap="round" stroke-linejoin="round" d="M 144.04154,48.000675 A 15.999945,15.999945 0 0 1 130.18519,40.00069"/><path id="galaFolder6" stroke-linecap="round" stroke-linejoin="round" d="M 144.04154,48.000675 A 15.999945,15.999945 0 0 0 130.18519,56.00066"/><path id="galaFolder7" stroke-linecap="butt" stroke-linejoin="miter" stroke-opacity="1" d="m 120.24118,24.000757 9.94401,15.999933"/><path id="galaFolder8" stroke-linecap="round" stroke-linejoin="round" d="m 106.3848,80.000579 a 15.999945,15.999945 0 0 0 13.85638,-7.999986"/><path id="galaFolder9" stroke-linecap="butt" stroke-linejoin="miter" stroke-opacity="1" d="M 120.24118,72.000593 130.18519,56.00066"/><path id="galaFoldera" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="m 16,80 90.38472,5.79e-4"/><path id="galaFolderb" stroke-linecap="round" stroke-linejoin="round" d="m 32.041967,16.000772 c -8.836533,3e-6 -15.999992,7.167628 -15.999978,16.004161"/><path id="galaFolderc" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="M 32.041967,16.000772 H 106.38498"/></g>`), g.Group(children),
	)
}

func Globe(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaGlobe0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" stroke-width="16"><circle id="galaGlobe1" cx="127.999" cy="128" r="112"/><path id="galaGlobe2" d="m 192.0027,219.91074 c -40.13552,-10.60764 2.8991,-36.95009 -12.98135,-56.80128 -7.0348,-8.7939 -24.94292,-22.44625 -16.06406,-37.68626 14.14662,-24.2819 74.44445,-16.70043 77.04541,2.65434"/><path id="galaGlobe3" d="m 52.26403,210.5115 c 11.735422,-10.73393 23.101176,-21.25355 28.928591,-35.30877 4.896866,-11.81082 11.221497,-27.74331 3.816787,-38.16679 -6.795901,-9.56626 -28.036264,1.74355 -34.229477,-8.2235 -5.19823,-8.36579 -1.703207,-26.91578 6.536171,-28.815794 C 71.025394,96.835252 95.999454,115.70076 91.525264,97.954854 71.651098,54.251844 101.47196,64.000005 79.998805,26.807419"/><path id="galaGlobe4" d="M 229.90244,81.526574 C 222.59641,76.902443 210.40368,72.548 197.29726,73.264682 184.1909,73.981363 156.7628,89.642259 150.64024,75.329052 143.20598,57.949662 165.8904,33.900895 191.58021,36.089261"/></g>`), g.Group(children),
	)
}

func Help(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaHelp0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-width="16"><circle id="galaHelp1" cx="181.018" cy="-.001" r="112.049" stroke-opacity="1" transform="rotate(45)"/><circle id="galaHelp2" cx="181.018" cy="-.001" r="48.021" transform="rotate(45)"/><path id="galaHelp3" stroke-opacity="1" d="M 171.33275,107.30371 226.85481,75.248027"/><path id="galaHelp4" stroke-opacity="1" d="M 148.69466,84.66561 180.75035,29.143566"/><path id="galaHelp5" stroke-opacity="1" d="M 84.667256,148.69302 29.145241,180.74869"/><path id="galaHelp6" stroke-opacity="1" d="M 107.30535,171.33111 75.249699,226.85316"/><path id="galaHelp7" stroke-opacity="1" d="m 148.69466,171.33111 32.05569,55.52205"/><path id="galaHelp8" stroke-opacity="1" d="m 171.33275,148.69302 55.52206,32.05567"/><path id="galaHelp9" stroke-opacity="1" d="M 107.30535,84.665618 75.249699,29.143573"/><path id="galaHelpa" stroke-opacity="1" d="M 84.667256,107.30372 29.145241,75.248046"/></g>`), g.Group(children),
	)
}

func Image(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaImage0"><g id="galaImage1" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-width="4.233" transform="scale(3.77953)"><rect id="galaImage2" width="59.267" height="59.267" x="4.233" y="4.233" stroke-opacity="1" ry="8.467"/><path id="galaImage3" stroke-opacity="1" d="m 25.4,33.866667 c -12.7,0 -16.9333334,8.466665 -21.1666669,8.466665"/><path id="galaImage4" stroke-opacity="1" d="m 25.4,33.866667 c 12.7,0 17.582797,8.466665 25.399999,8.466665"/><path id="galaImage5" stroke-opacity="1" d="M 63.5,38.099999 C 48.683332,38.099998 46.566666,50.8 38.1,50.8"/><circle id="galaImage6" cx="46.567" cy="21.167" r="8.467"/></g></g>`), g.Group(children),
	)
}

func Issue(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaIssue0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" stroke-width="16"><circle id="galaIssue1" cx="128" cy="128" r="40"/><circle id="galaIssue2" cx="128" cy="128" r="112"/></g>`), g.Group(children),
	)
}

func Layer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaLayer0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" stroke-width="16"><path id="galaLayer1" d="M 16,80 127.94695,15.974538"/><path id="galaLayer2" d="m 16,80 112,64"/><path id="galaLayer3" d="m 16,176 112,64"/><path id="galaLayer4" d="M 128,16 240,80"/><path id="galaLayer5" d="M 128,144 240,80"/><path id="galaLayer6" d="M 128,240 240,176"/><path id="galaLayer7" d="m 16,128 112,64"/><path id="galaLayer8" d="M 128,192 239.94695,128.0019"/><path id="galaLayer9" d="M 16,128 58.031909,104.01298"/><path id="galaLayera" d="M 16,176 58.032661,151.99127"/><path id="galaLayerb" d="m 239.94694,128.00191 -41.9796,-23.98708"/><path id="galaLayerc" d="M 240,176 197.96608,151.99056"/></g>`), g.Group(children),
	)
}

func Lock(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaLock0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-width="16"><path id="galaLock1" stroke-opacity="1" d="M 48.003124,207.99947 V 144.00668"/><path id="galaLock2" stroke-opacity="1" d="M 208.00205,207.99948 V 143.99989"/><path id="galaLock3" stroke-opacity="1" d="M 80.002888,239.99926 H 176.00225"/><path id="galaLock4" stroke-opacity="1" d="m 208.00205,207.99948 a 31.999787,31.999787 0 0 1 -31.9998,31.99978"/><path id="galaLock5" stroke-opacity="1" d="m 48.003124,207.99948 a 31.999787,31.999787 0 0 0 31.999764,31.99978"/><path id="galaLock6" stroke-opacity="1" d="m 128.00258,207.99949 v -31.9998"/><path id="galaLock7" d="M 47.996095,144.00668 A 15.999894,15.999894 0 0 1 63.995976,128.0068"/><path id="galaLock8" d="M 208.00205,143.99989 A 15.999894,15.999894 0 0 0 192.00218,128"/><path id="galaLock9" d="M 176.00225,64.00042 A 47.999683,47.999683 0 0 0 128.00258,16.00074"/><path id="galaLocka" d="M 80.002812,64.00042 A 47.999683,47.999683 0 0 1 128.00249,16.00074"/><path id="galaLockb" stroke-opacity="1" d="M 80.002888,128 V 63.907475"/><path id="galaLockc" stroke-opacity="1" d="M 176.00225,64.000424 V 128"/><path id="galaLockd" stroke-opacity="1" d="M 64.003006,128 H 192.00218"/></g>`), g.Group(children),
	)
}

func Mouse(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaMouse0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-width="16"><path id="galaMouse1" stroke-opacity="1" d="m 63.998705,64.001121 c 0,64.001289 -16.00032,48.000969 -16.00032,96.001929"/><path id="galaMouse2" stroke-opacity="1" d="m 192.00129,64.001121 c 0,-16.000322 0,-22.762857 -16.00032,-32.000645"/><path id="galaMouse3" stroke-opacity="1" d="m 128,48.000796 v 32.00065"/><path id="galaMouse4" stroke-opacity="1" d="m 79.99903,32.000476 c 0,0 32.00064,-16.000321 48.00097,-16.000321 16.00032,0 48.00097,16.000321 48.00097,16.000321"/><path id="galaMouse5" stroke-opacity="1" d="m 192.00129,64.001121 c 0,64.001289 16.00032,48.000969 16.00032,96.001929"/><path id="galaMouse6" d="m 208.00161,160.00305 a 80.00161,80.00161 0 0 1 -40.00081,69.28343 80.00161,80.00161 0 0 1 -80.001607,0 80.00161,80.00161 0 0 1 -40.000803,-69.28343"/><path id="galaMouse7" stroke-opacity="1" d="m 63.998705,64.001125 c 0,-16.000326 2e-6,-22.762857 16.000325,-32.000649"/></g>`), g.Group(children),
	)
}

func Multi(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaMulti0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" stroke-width="16"><circle id="galaMulti1" cx="56" cy="56" r="40"/><circle id="galaMulti2" cx="200" cy="56" r="40"/><circle id="galaMulti3" cx="56" cy="200" r="40"/><circle id="galaMulti4" cx="200" cy="200" r="40"/></g>`), g.Group(children),
	)
}

func Orbit(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaOrbit0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" stroke-width="16"><circle id="galaOrbit1" cx="208" cy="48.004" r="32.004"/><path id="galaOrbit2" d="M 226.27213,74.276705 A 112.01361,112.01361 0 0 1 195.13671,217.67055 112.01361,112.01361 0 0 1 48.774885,207.20976 112.01361,112.01361 0 0 1 38.349122,60.84543 112.01361,112.01361 0 0 1 181.75042,29.744337"/><path id="galaOrbit3" d="M 219.63599,191.62385 A 112.01361,112.01361 0 0 0 127.86254,144.01566 112.01361,112.01361 0 0 0 36.197622,191.83249"/></g>`), g.Group(children),
	)
}

func PortraitOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaPortrait10" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-width="16"><path id="galaPortrait11" stroke-opacity="1" d="M 16.056766,240 H 239.94325"/><path id="galaPortrait12" stroke-opacity="1" d="m 16.056766,240 c 0,-47.97569 31.983762,-64.0813 31.983762,-64.0813 15.991919,-9.23291 31.9838,0 47.975681,-15.99189"/><path id="galaPortrait13" stroke-opacity="1" d="m 96.016209,159.92681 c 0,0 15.991921,15.99189 31.983801,15.99189"/><path id="galaPortrait14" stroke-opacity="1" d="m 239.94325,240 c 0,-47.97569 -31.9838,-64.0813 -31.9838,-64.0813 -15.99188,-9.23291 -31.98376,0 -47.97564,-15.99189"/><path id="galaPortrait15" stroke-opacity="1" d="m 159.98381,159.92681 c 0,0 -15.99192,15.99189 -31.9838,15.99189"/><path id="galaPortrait16" d="m 128.00001,15.99977 c 31.9838,0 47.97568,15.991881 47.97568,63.967561 0,39.979759 -31.9838,63.967599 -47.97568,63.967599"/><path id="galaPortrait17" d="m 128.00001,15.99977 c -31.983801,0 -47.975682,15.991881 -47.975682,63.967561 0,39.979759 31.983802,63.967599 47.975682,63.967599"/><path id="galaPortrait18" stroke-opacity="1" d="m 159.98381,159.92681 -7.41232,-27.66304"/><path id="galaPortrait19" stroke-opacity="1" d="m 96.016209,159.92681 7.412321,-27.66304"/></g>`), g.Group(children),
	)
}

func PortraitTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaPortrait20" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-width="16"><path id="galaPortrait21" stroke-opacity="1" d="m 49.616696,208.00014 c 0,0 30.383521,0 46.383451,-15.99992"/><path id="galaPortrait22" stroke-opacity="1" d="m 96.000147,192.00022 c 0,0 15.999943,15.99992 31.999853,15.99992"/><path id="galaPortrait23" stroke-opacity="1" d="m 206.38331,208.00012 c 0,0 -30.38351,2e-5 -46.38345,-15.9999"/><path id="galaPortrait24" stroke-opacity="1" d="m 159.99986,192.00022 c 0,0 -15.99994,15.99992 -31.99986,15.99992"/><path id="galaPortrait25" d="m 128,48.000858 c 31.99986,0 47.9998,15.999926 47.9998,63.999712 -1e-5,39.99983 -31.99988,63.99971 -47.9998,63.99971"/><path id="galaPortrait26" d="m 128,48.000858 c -31.999857,0 -47.999787,15.999926 -47.999783,63.999712 3e-6,39.99983 31.999873,63.99971 47.999793,63.99971"/><path id="galaPortrait27" stroke-opacity="1" d="m 159.99986,192.00022 -7.41601,-27.67694"/><path id="galaPortrait28" stroke-opacity="1" d="m 96.000147,192.00022 7.416013,-27.67694"/><circle id="galaPortrait29" cx="128" cy="128" r="111.999"/></g>`), g.Group(children),
	)
}

func Radar(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaRadar0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-width="16"><path id="galaRadar1" d="M 224.97499,72.065141 A 111.96707,111.96707 0 0 1 207.18133,207.22135 111.96707,111.96707 0 0 1 72.025105,225.01501 111.96707,111.96707 0 0 1 19.856764,99.069471 111.96707,111.96707 0 0 1 128.00866,16.081608"/><path id="galaRadar2" stroke-opacity="1" d="M 128.00866,16.081615 V 112.05339"/><path id="galaRadar3" d="m 144.00394,128.04868 a 15.995295,15.995295 0 0 1 -15.99528,15.9953 15.995295,15.995295 0 0 1 -15.99529,-15.9953 15.995295,15.995295 0 0 1 15.99529,-15.99529 15.995295,15.995295 0 0 1 15.99528,15.99529 z"/><path id="galaRadar4" d="M 183.41795,96.058104 A 63.981179,63.981179 0 0 1 173.25016,173.29022 63.981179,63.981179 0 0 1 96.018049,183.45803 63.981179,63.981179 0 0 1 66.207591,111.48915 63.981179,63.981179 0 0 1 128.00866,64.067515"/></g>`), g.Group(children),
	)
}

func Remove(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaRemove0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" stroke-width="16"><circle id="galaRemove1" cx="128" cy="128" r="112"/><path id="galaRemove2" d="M 80.000004,128 H 176.00001"/></g>`), g.Group(children),
	)
}

func Search(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaSearch0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-miterlimit="4" stroke-width="16"><path id="galaSearch1" stroke-linecap="butt" stroke-linejoin="miter" stroke-opacity="1" d="m 89.074145,145.23139 -68.17345,68.17344"/><path id="galaSearch2" stroke-linecap="butt" stroke-linejoin="miter" stroke-opacity="1" d="M 111.27275,167.42999 43.099304,235.60344"/><path id="galaSearch3" stroke-linecap="round" stroke-linejoin="round" d="m 43.099305,235.60344 a 15.696788,15.696788 0 0 1 -22.19861,0"/><path id="galaSearch4" stroke-linecap="round" stroke-linejoin="round" d="m 20.900695,213.40483 a 15.696788,15.696788 0 0 0 0,22.19861"/><path id="galaSearch5" stroke-linecap="round" stroke-linejoin="round" d="M 240.65575,86.483932 A 70.635544,70.635544 0 0 1 170.0202,157.11948 70.635544,70.635544 0 0 1 99.384659,86.483932 70.635544,70.635544 0 0 1 170.0202,15.848389 70.635544,70.635544 0 0 1 240.65575,86.483932 Z"/><path id="galaSearch6" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="m 89.074145,145.23139 22.198605,22.1986"/><path id="galaSearch7" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="m 100.17344,156.33068 19.89988,-19.89987"/><path id="galaSearch8" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="m 70.126446,164.17908 22.198606,22.1986"/><path id="galaSearch9" stroke-linecap="round" stroke-linejoin="round" d="M 209.26216,86.483936 A 39.241967,39.241967 0 0 1 170.0202,125.7259"/></g>`), g.Group(children),
	)
}

func Secure(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaSecure0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" stroke-width="16"><path id="galaSecure1" d="m 127.99999,239.96468 c 0,0 95.98506,-31.99503 95.98506,-111.98257"/><path id="galaSecure2" d="M 223.98505,127.98211 V 31.997059 c 0,0 -31.99502,-15.997511 -95.98506,-15.997511"/><path id="galaSecure3" d="m 128,239.96468 c 0,0 -95.985056,-31.99503 -95.985056,-111.98257"/><path id="galaSecure4" d="M 32.014944,127.98211 V 31.997059 c 0,0 31.995019,-15.997509 95.985056,-15.997509"/><path id="galaSecure5" d="M 191.99003,63.99208 C 128,111.9846 112.00249,175.97464 112.00249,175.97464 c 0,0 -15.997511,-19.0946 -31.995019,-31.99502"/></g>`), g.Group(children),
	)
}

func Select(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaSelect0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-width="16"><path id="galaSelect1" d="M 16.000736,48.000563 A 31.999783,31.999783 0 0 1 48.000519,16.000782"/><path id="galaSelect2" d="m -239.99926,48.000563 a 31.999783,31.999783 0 0 1 31.99978,-31.999781" transform="scale(-1 1)"/><path id="galaSelect3" d="m -239.99926,-207.99947 a 31.999783,31.999783 0 0 1 31.99978,-31.99978" transform="scale(-1)"/><path id="galaSelect4" d="m 16.000736,-207.99947 a 31.999783,31.999783 0 0 1 31.999783,-31.99978" transform="scale(1 -1)"/><path id="galaSelect5" stroke-opacity="1" d="m 239.99923,143.99987 v 31.9998"/><path id="galaSelect6" stroke-opacity="1" d="M 239.99923,80.000312 V 112.00011"/><path id="galaSelect7" stroke-opacity="1" d="m 16.000747,143.99991 v 31.99976"/><path id="galaSelect8" stroke-opacity="1" d="M 16.000751,80.000312 V 112.00011"/><path id="galaSelect9" stroke-opacity="1" d="M 112.00008,16.000744 H 80.000284"/><path id="galaSelecta" stroke-opacity="1" d="M 175.99964,16.000748 H 143.99988"/><path id="galaSelectb" stroke-opacity="1" d="M 112.00008,239.99922 H 80.000284"/><path id="galaSelectc" stroke-opacity="1" d="M 175.99964,239.99922 H 143.99988"/><path id="galaSelectd" stroke-opacity="1" d="M 96.000202,127.99999 H 159.99976"/><path id="galaSelecte" stroke-opacity="1" d="M 128,96.000192 V 159.99979"/></g>`), g.Group(children),
	)
}

func Settings(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaSettings0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" stroke-width="16"><path id="galaSettings1" d="M 48.000002,16 H 208 c 17.728,0 32,14.272 32,32 v 160 c 0,17.728 -14.272,32 -32,32 H 48.000002 c -17.728,0 -32,-14.272 -32,-32 V 48 c 0,-17.728 14.272,-32 32,-32 z"/><path id="galaSettings2" d="M 64.000006,64.000001 H 79.999993"/><path id="galaSettings3" d="m 79.999996,-96.000015 a 16,16 0 0 1 -16,16 16,16 0 0 1 -16,-16 16,16 0 0 1 16,-16.000005 16,16 0 0 1 16,16.000005 z" transform="rotate(90)"/><path id="galaSettings4" d="m 112.00001,64.000353 79.99997,-3.52e-4"/><path id="galaSettings5" d="M 191.99998,128 H 176"/><path id="galaSettings6" d="m 144,159.99997 a 16,16 0 0 1 -16,16 16,16 0 0 1 -16,-16 16,16 0 0 1 16,-16 16,16 0 0 1 16,16 z" transform="matrix(0 1 1 0 0 0)"/><path id="galaSettings7" d="M 143.99998,128.00035 64.000006,128"/><path id="galaSettings8" d="M 64.000006,192.00001 H 79.999993"/><path id="galaSettings9" d="m 208,-96.000015 a 16,16 0 0 1 -16,16 16,16 0 0 1 -16,-16 16,16 0 0 1 16,-16.000005 16,16 0 0 1 16,16.000005 z" transform="rotate(90)"/><path id="galaSettingsa" d="m 112.00001,192.00036 79.99997,-3.5e-4"/></g>`), g.Group(children),
	)
}

func Shield(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaShield0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" stroke-width="16"><path id="galaShield1" d="m 128,239.89376 c 0,0 -95.908938,-31.96965 -95.908938,-111.89376"/><path id="galaShield2" d="M 32.091062,128 V 32.09106 c 0,0 31.969645,-15.984825 95.908938,-15.984825"/><path id="galaShield3" d="m 128,239.89376 c 0,0 95.90894,-31.96965 95.90894,-111.89376"/><path id="galaShield4" d="M 223.90894,128 V 32.09106 c 0,0 -31.96965,-15.984823 -95.90894,-15.984823"/><path id="galaShield5" d="m 128,239.89376 2.8e-4,-223.787523"/></g>`), g.Group(children),
	)
}

func SidebarLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaSidebarLeft0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" stroke-width="16"><path id="galaSidebarLeft1" d="m 16,64 224.93778,0.09256"/><path id="galaSidebarLeft2" d="M 48,16 H 207.91114 C 225.62929,16 240,30.281849 240,48 v 160 c 0,17.71816 -14.28185,32 -32,32 H 48 C 30.281848,240 16.069099,225.73073 16.06221,208.01257 L 16,48 C 15.993112,30.281851 30.281848,16 48,16 Z"/><path id="galaSidebarLeft3" d="M 96,64 V 240"/></g>`), g.Group(children),
	)
}

func SidebarRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaSidebarRight0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" stroke-width="16"><path id="galaSidebarRight1" d="m 16,64 224.93778,0.09256"/><path id="galaSidebarRight2" d="M 48,16 H 207.91114 C 225.62929,16 240,30.281849 240,48 v 160 c 0,17.71816 -14.28185,32 -32,32 H 48 C 30.281848,240 16.069099,225.73073 16.06221,208.01257 L 16,48 C 15.993112,30.281851 30.281848,16 48,16 Z"/><path id="galaSidebarRight3" d="M 160,64 V 240"/></g>`), g.Group(children),
	)
}

func Store(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaStore0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-miterlimit="4" stroke-width="16"><path id="galaStore1" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="M 48.000545,207.99268 V 128"/><path id="galaStore2" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="M 207.99947,207.99946 V 127.8564"/><path id="galaStore3" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="M 32.000626,16.00074 H 223.99935"/><path id="galaStore4" stroke-linecap="round" stroke-linejoin="round" d="m 183.99963,100.00018 a 27.999813,27.999813 0 0 1 -13.99991,24.24855 27.999813,27.999813 0 0 1 -27.99981,0 A 27.999813,27.999813 0 0 1 128,100.00018"/><path id="galaStore5" stroke-linecap="round" stroke-linejoin="round" d="m 239.99926,100.00018 a 27.999813,27.999813 0 0 1 -13.9999,24.24855 27.999813,27.999813 0 0 1 -27.99982,0 27.999813,27.999813 0 0 1 -13.9999,-24.24855"/><path id="galaStore6" stroke-linecap="round" stroke-linejoin="round" d="m 72.000366,100.00067 a 27.999813,27.999813 0 0 1 -13.999907,24.24855 27.999813,27.999813 0 0 1 -27.999813,0 27.999813,27.999813 0 0 1 -13.999906,-24.24855"/><path id="galaStore7" stroke-linecap="round" stroke-linejoin="round" d="m 128,100.00067 a 27.999813,27.999813 0 0 1 -13.99991,24.24855 27.999813,27.999813 0 0 1 -27.999814,0 27.999813,27.999813 0 0 1 -13.999906,-24.24855"/><path id="galaStore8" stroke-linecap="butt" stroke-linejoin="miter" stroke-opacity="1" d="M 16.000734,100.00252 32.000626,16.00074"/><path id="galaStore9" stroke-linecap="butt" stroke-linejoin="miter" stroke-opacity="1" d="m 223.99935,16.00074 15.99989,83.99945"/><path id="galaStorea" stroke-linecap="butt" stroke-linejoin="miter" stroke-opacity="1" d="M 183.99963,100.00108 175.99967,16.00074"/><path id="galaStoreb" stroke-linecap="butt" stroke-linejoin="miter" stroke-opacity="1" d="M 72.000386,100.00067 80.000307,16.00074"/><path id="galaStorec" stroke-linecap="butt" stroke-linejoin="miter" stroke-opacity="1" d="M 127.99999,100.00093 V 16.00074"/><path id="galaStored" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="M 32.000626,191.99958 H 223.99935"/><path id="galaStoree" stroke-linecap="butt" stroke-linejoin="miter" stroke-opacity="1" d="M 80.000307,239.99926 H 175.99967"/><path id="galaStoref" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="m 207.99944,207.99947 a 31.999786,31.999786 0 0 1 -31.99979,31.99978"/><path id="galaStoreg" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="M -48.000534,207.99947 A 31.999786,31.999786 0 0 1 -80.00032,239.99925" transform="scale(-1 1)"/></g>`), g.Group(children),
	)
}

func Terminal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaTerminal0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" stroke-width="16"><path id="galaTerminal1" d="m 16,64 224.93778,0.09256"/><path id="galaTerminal2" d="M 48,16 H 207.91114 C 225.62929,16 240,30.281849 240,48 v 160 c 0,17.71816 -14.28185,32 -32,32 H 48 C 30.281848,240 16.069099,225.73073 16.06221,208.01257 L 16,48 C 15.993112,30.281851 30.281848,16 48,16 Z"/><path id="galaTerminal3" d="M 191.96444,64.092555 192,16"/><path id="galaTerminal4" d="M 48,208 80,176 48,144"/><path id="galaTerminal5" d="m 96,208 h 32"/></g>`), g.Group(children),
	)
}

func Tv(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaTv0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" stroke-width="16"><path id="galaTv1" d="M 191.99999,208 H 64 C 46.272,208 32,193.728 32,176 V 96 C 32,78.272 46.272,64 64,64 h 128 c 17.728,0 32,14.272 32,32 v 80 c 0,17.728 -14.272,32 -32,32"/><path id="galaTv2" d="m 96,240 h 64"/><path id="galaTv3" d="M 128,64 176,16"/><path id="galaTv4" d="M 128,64 80,16"/></g>`), g.Group(children),
	)
}

func Unlock(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaUnlock0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-width="16"><path id="galaUnlock1" stroke-opacity="1" d="M 48.003124,207.99947 V 144.00668"/><path id="galaUnlock2" stroke-opacity="1" d="M 208.00205,207.99948 V 143.99989"/><path id="galaUnlock3" stroke-opacity="1" d="M 80.002888,239.99926 H 176.00225"/><path id="galaUnlock4" stroke-opacity="1" d="m 208.00205,207.99948 a 31.999787,31.999787 0 0 1 -31.9998,31.99978"/><path id="galaUnlock5" stroke-opacity="1" d="m 48.003124,207.99948 a 31.999787,31.999787 0 0 0 31.999764,31.99978"/><path id="galaUnlock6" stroke-opacity="1" d="m 128.00258,207.99949 v -31.9998"/><path id="galaUnlock7" d="M 47.996095,144.00668 A 15.999894,15.999894 0 0 1 63.995976,128.0068"/><path id="galaUnlock8" d="M 208.00205,143.99989 A 15.999894,15.999894 0 0 0 192.00218,128"/><path id="galaUnlock9" d="M 176.00225,64.00042 A 47.999683,47.999683 0 0 0 128.00258,16.00074"/><path id="galaUnlocka" d="M 80.002812,64.00042 A 47.999683,47.999683 0 0 1 128.00249,16.00074"/><path id="galaUnlockb" stroke-opacity="1" d="M 80.002888,128 V 63.907475"/><path id="galaUnlockc" stroke-opacity="1" d="m 176.00225,64.000424 0,15.999576"/><path id="galaUnlockd" stroke-opacity="1" d="M 64.003006,128 H 192.00218"/></g>`), g.Group(children),
	)
}

func Usb(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaUsb0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-width="16"><path id="galaUsb1" stroke-opacity="1" d="M 48.042103,208.01896 V 112.07317"/><path id="galaUsb2" stroke-opacity="1" d="M 207.96306,208.01897 V 112.06638"/><path id="galaUsb3" stroke-opacity="1" d="M 80.026346,240.00316 H 175.97888"/><path id="galaUsb4" stroke-opacity="1" d="m 207.96306,208.01897 a 31.98419,31.98419 0 0 1 -31.98418,31.98419"/><path id="galaUsb5" stroke-opacity="1" d="m 48.042103,208.01897 a 31.98419,31.98419 0 0 0 31.984243,31.98419"/><path id="galaUsb6" d="M 48.035077,112.07316 A 15.992095,15.992095 0 0 1 64.027161,96.081066"/><path id="galaUsb7" d="M 207.96306,112.06637 A 15.992095,15.992095 0 0 0 191.97096,96.074276"/><path id="galaUsb8" stroke-opacity="1" d="M 64.034224,96.074284 H 191.97096"/><path id="galaUsb9" d="M 80.026346,32.112684 C 80.026347,23.280513 87.167829,16.000027 96,16"/><path id="galaUsba" d="M 175.97888,32.105897 C 175.97888,23.273697 168.8322,15.999986 160,16"/><path id="galaUsbb" stroke-opacity="1" d="M 96.018393,16 H 160"/><path id="galaUsbc" stroke-opacity="1" d="M 80.026346,32.105904 V 96.074281"/><path id="galaUsbd" stroke-opacity="1" d="M 175.97888,32.105904 V 96.074281"/><path id="galaUsbe" stroke-opacity="1" d="m 112.01048,16.113809 0,31.984193"/><path id="galaUsbf" stroke-opacity="1" d="M 143.99468,16.113809 V 48.098002"/></g>`), g.Group(children),
	)
}

func Video(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaVideo0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" stroke-width="16"><path id="galaVideo1" d="M 47.999992,48 H 160 c 17.728,0 32,14.272 32,32 v 96 c 0,17.728 -14.272,32 -32,32 H 47.999992 c -17.728,0 -32,-14.272 -32,-32 V 80 c 0,-17.728 14.272,-32 32,-32 z"/><path id="galaVideo2" d="m 192,160 48,32 V 64 l -48,32"/></g>`), g.Group(children),
	)
}

func Window(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g id="galaWindow0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" stroke-width="16"><path id="galaWindow1" d="m 16,64 224.93778,0.09256"/><path id="galaWindow2" d="M 48,16 H 207.91114 C 225.62929,16 240,30.281849 240,48 v 160 c 0,17.71816 -14.28185,32 -32,32 H 48 C 30.281848,240 16.069099,225.73073 16.06221,208.01257 L 16,48 C 15.993112,30.281851 30.281848,16 48,16 Z"/><path id="galaWindow3" d="M 191.96444,64.092555 192,16"/></g>`), g.Group(children),
	)
}
