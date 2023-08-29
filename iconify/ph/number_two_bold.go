package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberTwoBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M180 208a12 12 0 0 1-12 12H88a12 12 0 0 1-9.6-19.2l71.95-95.92a28 28 0 1 0-48-28.06a12 12 0 0 1-22-9.62a52.32 52.32 0 0 1 6.13-10.49a52 52 0 0 1 83.06 62.59L112 196h56a12 12 0 0 1 12 12Z"/>`),
		g.Group(children),
	)
}