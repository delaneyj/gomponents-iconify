package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalCellularAddOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.7 20.3L20.275 3.725q.475-.475 1.088-.212t.612.937v8.2q-.45-.275-.95-.437t-1.05-.263v-5.1L6.825 20H13.2q.2.575.5 1.075t.675.925h-9.95q-.675 0-.937-.613T3.7 20.3ZM18 19h-2q-.425 0-.713-.288T15 18q0-.425.288-.713T16 17h2v-2q0-.425.288-.713T19 14q.425 0 .713.288T20 15v2h2q.425 0 .713.288T23 18q0 .425-.288.713T22 19h-2v2q0 .425-.288.713T19 22q-.425 0-.713-.288T18 21v-2ZM6.825 20l13.15-13.15l-3.413 3.413l-3.025 3.024l-3.087 3.088L6.825 20Z"/>`),
		g.Group(children),
	)
}