package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarDot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCalendarDot0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="40" height="40" x="4" y="4" fill="#fff" stroke="#fff" rx="2"/><path stroke="#000" d="M4 14h40"/><path stroke="#fff" d="M44 11v12"/><path stroke="#000" d="M12 22h4m6 0h4m6 0h4m-24 7h4m6 0h4m6 0h4m-24 7h4m6 0h4m6 0h4"/><path stroke="#fff" d="M4 11v12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCalendarDot0)"/>`),
		g.Group(children),
	)
}