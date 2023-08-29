package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Highlight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h3.5v2H4v1.5H2V2Zm5.5 0h3v2h-3V2Zm5 0h3v2h-3V2Zm5 0H21v3.5h-2V4h-1.5V2ZM4 7.5v3H2v-3h2Zm17 0v3h-2v-3h2Zm-17 5v3H2v-3h2ZM20 14v2h-2.586l4 4L20 21.414l-4-4V20h-2v-6h6ZM4 17.5V19h1.5v2H2v-3.5h2ZM7.5 19h3v2h-3v-2Z"/>`),
		g.Group(children),
	)
}