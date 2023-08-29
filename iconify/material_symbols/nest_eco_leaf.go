package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestEcoLeaf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.3 19.7q-.275-.275-.275-.7t.275-.7l1.375-1.375q-.8-1.025-1.238-2.275T4 12q0-3.35 2.325-5.675T12 4h8v8q0 3.35-2.325 5.675T12 20q-1.4 0-2.638-.438T7.1 18.325L5.7 19.7q-.275.275-.7.275t-.7-.275Zm4.05-4.05q.275.325.7.313t.725-.313l3.925-3.925q.3-.3.3-.712t-.3-.713q-.275-.275-.7-.275t-.7.275l-3.95 3.95q-.275.275-.275.688t.275.712Z"/>`),
		g.Group(children),
	)
}