package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestEcoLeafOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 20q-1.4 0-2.638-.438T7.1 18.325L5.7 19.7q-.275.275-.7.275t-.7-.275q-.275-.275-.275-.7t.275-.7l1.375-1.375q-.8-1.025-1.238-2.275T4 12q0-3.35 2.325-5.675T12 4h8v8q0 3.35-2.325 5.675T12 20Zm0-2q2.5 0 4.25-1.75T18 12V6h-6Q9.5 6 7.75 7.75T6 12q0 .975.3 1.863t.825 1.612L12.3 10.3q.275-.275.7-.275t.7.275q.3.3.3.713t-.3.712L8.525 16.9q.725.525 1.613.813T12 18Zm0-6Z"/>`),
		g.Group(children),
	)
}