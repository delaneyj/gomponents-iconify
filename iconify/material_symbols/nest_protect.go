package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestProtect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 21q-2.1 0-3.55-1.45T3 16V8q0-2.1 1.45-3.55T8 3h8q2.1 0 3.55 1.45T21 8v8q0 2.1-1.45 3.55T16 21H8Zm0-2h8q1.275 0 2.138-.863T19 16V8q0-1.275-.863-2.138T16 5H8q-1.275 0-2.138.863T5 8v8q0 1.275.863 2.138T8 19Zm4-2q-2.1 0-3.55-1.45T7 12q0-2.1 1.45-3.55T12 7q2.1 0 3.55 1.45T17 12q0 2.1-1.45 3.55T12 17Zm0-2q1.275 0 2.138-.863T15 12q0-1.275-.863-2.138T12 9q-1.275 0-2.138.863T9 12q0 1.275.863 2.138T12 15Z"/>`),
		g.Group(children),
	)
}