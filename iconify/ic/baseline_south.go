package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaselineSouth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19 15l-1.41-1.41L13 18.17V2h-2v16.17l-4.59-4.59L5 15l7 7l7-7z"/>`),
		g.Group(children),
	)
}