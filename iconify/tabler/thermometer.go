package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Thermometer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 5a2.828 2.828 0 0 1 0 4l-8 8H7v-4l8-8a2.828 2.828 0 0 1 4 0zm-3 2l-1.5-1.5M13 10l-1.5-1.5M10 13l-1.5-1.5M7 17l-3 3"/>`),
		g.Group(children),
	)
}