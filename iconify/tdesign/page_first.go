package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageFirst(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.5 5v14h-2V5h2Zm9.164 1.75L12.414 12l5.25 5.25l-1.414 1.414L9.586 12l6.664-6.664l1.414 1.414Z"/>`),
		g.Group(children),
	)
}