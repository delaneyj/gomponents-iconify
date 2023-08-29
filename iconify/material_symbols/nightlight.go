package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nightlight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 22q-2.05 0-3.875-.788t-3.188-2.15q-1.362-1.362-2.15-3.187T4 12q0-2.075.788-3.888t2.15-3.174q1.362-1.363 3.187-2.15T14 2q1.35 0 2.625.35t2.375 1q-2.275 1.325-3.638 3.588T14 12q0 2.8 1.363 5.063T19 20.65q-1.1.65-2.375 1T14 22Z"/>`),
		g.Group(children),
	)
}