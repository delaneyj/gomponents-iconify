package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeftDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18.01 7.404l-8.19 8.192h6.364v2h-9.78V7.818h2v6.364l8.193-8.192l1.414 1.414Z"/>`),
		g.Group(children),
	)
}