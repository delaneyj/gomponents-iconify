package gala

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Calendar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g id="galaCalendar0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-miterlimit="4" stroke-width="16"><path id="galaCalendar1" stroke-linecap="butt" stroke-linejoin="miter" stroke-opacity="1" d="M 31.999978,31.999961 H 224.00004"/><path id="galaCalendar2" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="m 15.999975,47.999965 -3e-6,176.000055"/><path id="galaCalendar3" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="M 240.00002,47.999965 V 224.00002"/><path id="galaCalendar4" stroke-linecap="round" stroke-linejoin="round" d="m 31.999978,31.999961 c -8.836576,0 -16.000003,7.163446 -16.000003,16.000004"/><path id="galaCalendar5" stroke-linecap="round" stroke-linejoin="round" d="m 224.00004,31.999961 c 8.83657,-4e-6 15.99998,7.163443 15.99998,16.000004"/><path id="galaCalendar6" stroke-linecap="butt" stroke-linejoin="miter" stroke-opacity="1" d="M 224.00004,240.00002 H 31.999978"/><path id="galaCalendar7" stroke-linecap="round" stroke-linejoin="round" d="m 224.00004,240.00002 a 16.000004,16.000004 0 0 0 15.99998,-16"/><path id="galaCalendar8" stroke-linecap="round" stroke-linejoin="round" d="m 31.999978,240.00002 a 16.000004,16.000004 0 0 1 -16.000011,-16"/><path id="galaCalendar9" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="M 128.00001,47.999965 V 15.999962"/><path id="galaCalendara" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="M 160.00003,47.999965 V 15.999962"/><path id="galaCalendarb" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="M 192.00002,47.999965 V 15.999962"/><path id="galaCalendarc" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="M 95.999985,47.999965 V 15.999962"/><path id="galaCalendard" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="M 64.000001,47.999965 V 15.999962"/><path id="galaCalendare" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="M 15.999975,95.999972 H 240.00002"/></g>`),
		g.Group(children),
	)
}