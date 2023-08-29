package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarDot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTCalendarDot0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="40" height="40" x="4" y="4" fill="#555" rx="2"/><path d="M4 14h40m0-3v12m-32-1h4m6 0h4m6 0h4m-24 7h4m6 0h4m6 0h4m-24 7h4m6 0h4m6 0h4M4 11v12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTCalendarDot0)"/>`),
		g.Group(children),
	)
}