package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarThirtyTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCalendarThirtyTwo0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M4 8h40v36H4z"/><path stroke="#000" stroke-linecap="round" d="M28 20v14h8V20h-8Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" d="M17 4v8m14-8v8"/><path stroke="#000" stroke-linecap="round" d="M12 20h8v14h-8m8-7h-6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCalendarThirtyTwo0)"/>`),
		g.Group(children),
	)
}