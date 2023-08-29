package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NightlightRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 22q-2.05 0-3.875-.788t-3.188-2.15q-1.362-1.362-2.15-3.187T4 12q0-2.075.788-3.888t2.15-3.174q1.362-1.363 3.187-2.15T14 2q1.1 0 2.125.238T18.1 2.9q.35.175.375.463t-.325.537q-1.9 1.375-3.025 3.475T14 12q0 2.525 1.125 4.625T18.15 20.1q.35.25.325.537t-.375.463q-.95.425-1.975.663T14 22Z"/>`),
		g.Group(children),
	)
}