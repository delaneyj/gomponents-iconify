package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarThirty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCalendarThirty0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="40" height="40" x="4" y="4" fill="#fff" stroke="#fff" rx="2"/><path stroke="#000" d="M4 14h40"/><path stroke="#fff" d="M4 11v12m40-12v12"/><path stroke="#000" d="M28 22v14h8V22h-8Z" clip-rule="evenodd"/><path stroke="#000" d="M12 22h8v14h-8m8-7h-6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCalendarThirty0)"/>`),
		g.Group(children),
	)
}