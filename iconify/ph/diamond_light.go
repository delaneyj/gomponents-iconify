package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiamondLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m233.92 118.14l-96.06-96.06a14 14 0 0 0-19.72 0l-96.06 96.06a14 14 0 0 0 0 19.72l96.06 96.06a14 14 0 0 0 19.72 0l96-96.06a13.94 13.94 0 0 0 0-19.72Zm-8.49 11.24l-96.05 96.06a2 2 0 0 1-2.76 0l-96.05-96.06a2 2 0 0 1 0-2.76l96.05-96.06a2 2 0 0 1 2.76 0l96.05 96.06a2 2 0 0 1 0 2.76Z"/>`),
		g.Group(children),
	)
}