package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileCodeOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V23H3V1Zm2 2v18h14V9h-6V3H5Zm10 .414V7h3.586L15 3.414Zm-3.793 8.793L9.414 14l1.793 1.793l-1.414 1.414L6.586 14l3.207-3.207l1.414 1.414Zm3-1.414L17.414 14l-3.207 3.207l-1.414-1.414L14.586 14l-1.793-1.793l1.414-1.414Z"/>`),
		g.Group(children),
	)
}