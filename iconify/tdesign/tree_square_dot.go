package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreeSquareDot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.5 2h7v7H13v2.333h6.5V15H22v7h-7v-7h2.5v-1.667h-11V15H9v7H2v-7h2.5v-3.667H11V9H8.5V2Zm5 5V4h-3v3h3ZM4 17v3h3v-3H4Zm13 0v3h3v-3h-3Z"/>`),
		g.Group(children),
	)
}