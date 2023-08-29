package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContactPhoneOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19 18l2-2l-1.5-2h-1.65q-.15-.45-.25-.963T17.5 12q0-.525.1-1.012t.25-.988h1.65L21 8l-2-2q-1.35 1.05-2.175 2.663T16 12q0 1.725.825 3.338T19 18ZM0 21V3h24v18H0Zm15.9-2H22V5H2v14h.1q1.05-1.875 2.9-2.938T9 15q2.15 0 4 1.063T15.9 19ZM9 14q1.25 0 2.125-.875T12 11q0-1.25-.875-2.125T9 8q-1.25 0-2.125.875T6 11q0 1.25.875 2.125T9 14Zm-4.45 5h8.9q-.85-.95-2.013-1.475T9 17q-1.275 0-2.425.525T4.55 19ZM9 12q-.425 0-.713-.288T8 11q0-.425.288-.713T9 10q.425 0 .713.288T10 11q0 .425-.288.713T9 12Zm3 0Z"/>`),
		g.Group(children),
	)
}