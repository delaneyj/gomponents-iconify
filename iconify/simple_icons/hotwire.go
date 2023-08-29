package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hotwire(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.764 6.917l-3.48.81L16.32 0L7.236 11.705l4.334-.854l-4.087 7.982l2.364-.532L7.456 24l7.51-8.111l-2.77.64l4.568-9.612z"/>`),
		g.Group(children),
	)
}