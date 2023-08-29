package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmClock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSAlarmClock0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M24 44.333c10.125 0 18.333-8.208 18.333-18.333c0-10.125-8.208-18.333-18.333-18.333C13.875 7.667 5.666 15.875 5.666 26c0 10.125 8.209 18.333 18.334 18.333Z"/><path stroke="#000" stroke-linecap="round" d="m23.76 15.354l-.002 11.008l7.773 7.773"/><path stroke="#fff" stroke-linecap="round" d="m4 9l7-5m33 5l-7-5"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSAlarmClock0)"/>`),
		g.Group(children),
	)
}