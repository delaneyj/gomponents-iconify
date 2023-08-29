package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccountBalanceRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 17q-.425 0-.713-.288T5 16v-5.025q0-.425.288-.7T6 10q.425 0 .713.288T7 11v5.025q0 .425-.288.7T6 17Zm6 0q-.425 0-.713-.288T11 16v-5.025q0-.425.288-.7T12 10q.425 0 .713.288T13 11v5.025q0 .425-.288.7T12 17Zm-9.025 4q-.425 0-.7-.288T2 20q0-.425.288-.713T3 19h18.025q.425 0 .7.288T22 20q0 .425-.288.713T21 21H2.975ZM18 17q-.425 0-.713-.288T17 16v-5.025q0-.425.288-.7T18 10q.425 0 .713.288T19 11v5.025q0 .425-.288.7T18 17ZM12.9 1.45l8.425 4.2q.325.175.5.487t.175.688q0 .5-.362.837T20.775 8H3.25q-.5 0-.875-.338T2 6.825q0-.35.163-.675t.512-.475L11.1 1.45q.425-.2.9-.2t.9.2Z"/>`),
		g.Group(children),
	)
}