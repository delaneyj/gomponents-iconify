package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronUpS(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 8.316l4.95 4.95l-1.415 1.414L12 11.144L8.464 14.68L7.05 13.265L12 8.316Z"/>`),
		g.Group(children),
	)
}