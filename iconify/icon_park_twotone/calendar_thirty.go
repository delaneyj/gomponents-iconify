package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarThirty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTCalendarThirty0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="40" height="40" x="4" y="4" fill="#555" rx="2"/><path d="M4 14h40M4 11v12m40-12v12"/><path d="M28 22v14h8V22h-8Z" clip-rule="evenodd"/><path d="M12 22h8v14h-8m8-7h-6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTCalendarThirty0)"/>`),
		g.Group(children),
	)
}