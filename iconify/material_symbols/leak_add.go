package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeakAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 14v-2q1.85 0 3.488-.713T9.35 9.35q1.225-1.225 1.938-2.863T12 3h2q0 2.275-.863 4.275t-2.362 3.5q-1.5 1.5-3.5 2.363T3 14Zm0-4V8q2.075 0 3.538-1.463T8 3h2q0 2.9-2.05 4.95T3 10Zm0-4V3h3q0 1.25-.875 2.125T3 6Zm7 15q0-2.275.863-4.275t2.362-3.5q1.5-1.5 3.5-2.362T21 10v2q-1.85 0-3.488.713T14.65 14.65q-1.225 1.225-1.938 2.863T12 21h-2Zm4 0q0-2.9 2.05-4.95T21 14v2q-2.075 0-3.538 1.463T16 21h-2Zm4 0q0-1.25.875-2.125T21 18v3h-3Z"/>`),
		g.Group(children),
	)
}