package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RhythmLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 9h2v12H2V9Zm6-6h2v18H8V3Zm6 9h2v9h-2v-9Zm6-6h2v15h-2V6Z"/>`),
		g.Group(children),
	)
}