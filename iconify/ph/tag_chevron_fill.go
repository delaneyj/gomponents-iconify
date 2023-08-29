package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TagChevronFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m237.3 136.88l-42.66 64a16 16 0 0 1-13.31 7.12H24a8 8 0 0 1-6.58-12.55L64 128L17.42 60.55A8 8 0 0 1 24 48h157.33a16 16 0 0 1 13.31 7.12l42.66 64a16 16 0 0 1 0 17.76Z"/>`),
		g.Group(children),
	)
}