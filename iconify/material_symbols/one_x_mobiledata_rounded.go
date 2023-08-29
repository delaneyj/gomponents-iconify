package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OneXMobiledataRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 9H5q-.425 0-.713-.288T4 8q0-.425.288-.713T5 7h2q.425 0 .713.288T8 8v8q0 .425-.288.713T7 17q-.425 0-.713-.288T6 16V9Zm8.65 4.65l-1.725 2.875q-.125.225-.338.35T12.1 17q-.575 0-.863-.5t.013-1l2.25-3.8l-1.925-3.175q-.3-.5-.013-1.013T12.425 7q.275 0 .513.138t.362.362l1.35 2.25l1.4-2.275q.125-.225.35-.35T16.9 7q.575 0 .863.5t-.013 1l-1.9 3.2l2.25 3.775q.3.5.013 1.012t-.888.513q-.275 0-.513-.138t-.362-.362l-1.7-2.85Z"/>`),
		g.Group(children),
	)
}