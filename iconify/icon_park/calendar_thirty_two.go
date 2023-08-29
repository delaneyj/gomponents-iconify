package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarThirtyTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="40" height="36" x="4" y="8" fill="#2F88FF" stroke="#000"/><path stroke="#fff" stroke-linecap="round" d="M28 20V34H36V20H28Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" d="M17 4V12"/><path stroke="#000" stroke-linecap="round" d="M31 4V12"/><path stroke="#fff" stroke-linecap="round" d="M12 20H20V34H12"/><path stroke="#fff" stroke-linecap="round" d="M20 27H14"/></g>`),
		g.Group(children),
	)
}