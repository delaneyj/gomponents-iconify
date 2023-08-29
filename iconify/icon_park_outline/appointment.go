package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Appointment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><circle cx="24" cy="11" r="7" stroke-linecap="round" stroke-linejoin="round"/><path stroke-linecap="round" stroke-linejoin="round" d="M4 41c0-8.837 8.059-16 18-16"/><circle cx="34" cy="34" r="9"/><path stroke-linecap="round" stroke-linejoin="round" d="M33 31v4h4"/></g>`),
		g.Group(children),
	)
}