package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 0v3h3v2h-3v3h-2V5h-3V3h3V0h2ZM2 2h11v2H4v16h16V10h2v12H2V2Zm11 6v10h-2V8h2Zm-5 3v7H6v-7h2Zm10 2v5h-2v-5h2Z"/>`),
		g.Group(children),
	)
}