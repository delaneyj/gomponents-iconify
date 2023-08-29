package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SystemRegulation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 1h22v17H1V1Zm2 2v13h18V3H3Zm6 2v5H7V5h2Zm4 0v3h-2V5h2Zm4 0v5h-2V5h2Zm-4 4v5h-2V9h2Zm-4 2v3H7v-3h2Zm8 0v3h-2v-3h2ZM3.222 21h17.556v2H3.222v-2Z"/>`),
		g.Group(children),
	)
}