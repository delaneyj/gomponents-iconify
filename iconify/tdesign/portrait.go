package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Portrait(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h7v2H4v5H2V2Zm13 0h7v7h-2V4h-5V2Zm-3 6.5a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Zm2.665 3.769a3.5 3.5 0 1 0-5.33 0A4.996 4.996 0 0 0 7 16.5v1h2v-1a3 3 0 1 1 6 0v1h2v-1a4.996 4.996 0 0 0-2.335-4.231ZM4 15v5h5v2H2v-7h2Zm18 0v7h-7v-2h5v-5h2Z"/>`),
		g.Group(children),
	)
}