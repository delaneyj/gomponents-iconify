package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Subscript(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 7l8 10m-8 0l8-10m8 13h-4l3.5-4a1.73 1.73 0 0 0-3.5-2"/>`),
		g.Group(children),
	)
}