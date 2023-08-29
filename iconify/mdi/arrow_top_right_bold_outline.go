package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTopRightBoldOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.64 4.22h14.14v14.14l-4.24-4.24l-5.66 5.66l-5.66-5.66l5.66-5.66l-4.24-4.24m12.02 2.12h-7.07l2.12 2.12l-5.66 5.66l2.83 2.83l5.66-5.66l2.12 2.12V6.34Z"/>`),
		g.Group(children),
	)
}