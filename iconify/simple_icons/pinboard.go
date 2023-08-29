package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pinboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13.352 14.585l-4.509 4.614l.72-4.062L3.428 7.57L0 7.753L7.58 0v2.953l7.214 6.646l4.513-1.105l-4.689 4.982L24 24l-10.648-9.415z"/>`),
		g.Group(children),
	)
}