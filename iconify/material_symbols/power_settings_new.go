package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerSettingsNew(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 13V3h2v10h-2Zm1 8q-1.85 0-3.488-.713T5.65 18.35q-1.225-1.225-1.938-2.863T3 12q0-2 .825-3.775T6.15 5.15L7.6 6.6q-1.25.95-1.925 2.375T5 12q0 2.9 2.05 4.95T12 19q2.925 0 4.963-2.05T19 12q0-1.6-.663-3.025T16.4 6.6l1.45-1.45q1.5 1.3 2.325 3.075T21 12q0 1.85-.713 3.488t-1.924 2.862q-1.213 1.225-2.85 1.938T12 21Z"/>`),
		g.Group(children),
	)
}