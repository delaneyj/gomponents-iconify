package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarThirty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="40" height="40" x="4" y="4" fill="#2F88FF" stroke="#000" rx="2"/><path stroke="#fff" d="M4 14H44"/><line x1="4" x2="4" y1="11" y2="23" stroke="#000"/><line x1="44" x2="44" y1="11" y2="23" stroke="#000"/><path stroke="#fff" d="M28 22V36H36V22H28Z" clip-rule="evenodd"/><path stroke="#fff" d="M12 22H20V36H12"/><path stroke="#fff" d="M20 29H14"/></g>`),
		g.Group(children),
	)
}