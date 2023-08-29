package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSplitVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 16v-3h-3v9h-2V2h2v9h3V8l4 4l-4 4M2 12l4 4v-3h3v9h2V2H9v9H6V8l-4 4Z"/>`),
		g.Group(children),
	)
}