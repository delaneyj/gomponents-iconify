package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HardDisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M44 29H4v13h40V29Z"/><path fill="currentColor" d="M35.5 38a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M4 29L9.038 4.999H39.02l4.98 24"/></g>`),
		g.Group(children),
	)
}