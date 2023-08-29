package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TaskAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 1h10v2h4v9h-2V5h-2v2H7V5H5v16h7v2H3V3h4V1Zm2 4h6V3H9v2Zm11 9v4h4v2h-4v4h-2v-4h-4v-2h4v-4h2Z"/>`),
		g.Group(children),
	)
}