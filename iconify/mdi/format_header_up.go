package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatHeaderUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M2 4h2v6h4V4h2v14H8v-6H4v6H2V4zm18.586 10.415L17 10.828l-3.586 3.587L12 13.001L17 8l5 5.001l-1.414 1.414z" fill="currentColor"/>`),
		g.Group(children),
	)
}