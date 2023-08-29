package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Verify(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 3h22v18H1V3Zm2 2v14h18V5H3Zm12.5 4.5a1 1 0 1 1 0 2a1 1 0 0 1 0-2Zm2.4 2.8a3 3 0 1 0-4.8 0a3.994 3.994 0 0 0-1.6 3.2v1h2v-1a2 2 0 1 1 4 0v1h2v-1c0-1.309-.628-2.47-1.6-3.2ZM5 9h4.5v2H5V9Zm0 4h4.5v2H5v-2Z"/>`),
		g.Group(children),
	)
}