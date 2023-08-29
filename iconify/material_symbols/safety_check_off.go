package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SafetyCheckOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.8 22.6l1.4-1.4L2.8 2.8L1.4 4.2L4 6.8v4.3q0 3.8 2.262 6.913T12 22q1.275-.35 2.438-1.012t2.112-1.638l3.25 3.25ZM12 17q-2.075 0-3.538-1.463T7 12q0-.5.088-.963t.262-.887l6.5 6.5q-.425.175-.887.263T12 17Zm6.85-.95q.575-1.175.863-2.425T20 11.1V5l-8-3l-5.2 1.95l3.375 3.4q.425-.175.875-.262T12 7q2.075 0 3.538 1.463T17 12q0 .5-.088.963t-.262.887l2.2 2.2Z"/>`),
		g.Group(children),
	)
}