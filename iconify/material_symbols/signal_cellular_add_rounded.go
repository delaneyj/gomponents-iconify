package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalCellularAddRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 22q-.425 0-.713-.288T18 21v-2h-2q-.425 0-.713-.288T15 18q0-.425.288-.713T16 17h2v-2q0-.425.288-.713T19 14q.425 0 .713.288T20 15v2h2q.425 0 .713.288T23 18q0 .425-.288.713T22 19h-2v2q0 .425-.288.713T19 22ZM4.425 22q-.675 0-.938-.613T3.7 20.3L20.275 3.725q.475-.475 1.088-.212t.612.937v8.2q-.65-.4-1.425-.587T19 11.875q-2.55 0-4.338 1.788T12.875 18q0 1.15.375 2.138T14.375 22h-9.95Z"/>`),
		g.Group(children),
	)
}