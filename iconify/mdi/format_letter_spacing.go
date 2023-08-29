package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatLetterSpacing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.4 10.7h4.1l-2-5.4l-2.1 5.4M15.6 3h1.7L22 15h-1.9l-1-2.6h-5.4l-1 2.6h-1.9l4.8-12m-4.4 0h1.9L8.4 15H6.7L2 3h1.9l3.6 9.7M19 22v-2H5v2l-3-3l3-3v2h14v-2l3 3l-3 3Z"/>`),
		g.Group(children),
	)
}