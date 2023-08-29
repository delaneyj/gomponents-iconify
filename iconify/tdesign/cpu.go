package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cpu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 .5V3h6V.5h2V3h4v4h2.5v2H21v6h2.5v2H21v4h-4v2.5h-2V21H9v2.5H7V21H3v-4H.5v-2H3V9H.5V7H3V3h4V.5h2ZM5 5v14h14V5H5Zm3 3h8v8H8V8Zm2 2v4h4v-4h-4Z"/>`),
		g.Group(children),
	)
}