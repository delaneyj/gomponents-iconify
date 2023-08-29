package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HdrWeak(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 16q-1.65 0-2.825-1.175T1 12q0-1.65 1.175-2.825T5 8q1.65 0 2.825 1.175T9 12q0 1.65-1.175 2.825T5 16Zm12 2q-2.5 0-4.25-1.75T11 12q0-2.5 1.75-4.25T17 6q2.5 0 4.25 1.75T23 12q0 2.5-1.75 4.25T17 18Zm0-2q1.65 0 2.825-1.175T21 12q0-1.65-1.175-2.825T17 8q-1.65 0-2.825 1.175T13 12q0 1.65 1.175 2.825T17 16Z"/>`),
		g.Group(children),
	)
}