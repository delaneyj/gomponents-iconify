package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HospitalBuilding(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22V7a1 1 0 0 1 1-1h4V2h10v4h4a1 1 0 0 1 1 1v15h-8v-5h-4v5H2M9 4v6h2V8h2v2h2V4h-2v2h-2V4H9M4 20h4v-3H4v3m0-5h4v-3H4v3m12 5h4v-3h-4v3m0-5h4v-3h-4v3m-6 0h4v-3h-4v3Z"/>`),
		g.Group(children),
	)
}