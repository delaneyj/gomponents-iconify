package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NatRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 15q-1.25 0-2.125-.875T1 12q0-1.25.875-2.125T4 9q.975 0 1.738.563T6.8 11H11v2H6.8q-.3.875-1.062 1.438T4 15Zm0-2q.425 0 .713-.288T5 12q0-.425-.288-.713T4 11q-.425 0-.713.288T3 12q0 .425.288.713T4 13Zm2 8.95q-.4.05-.7-.237T5 21q0-.4.288-.7T6 19.95q2.975-.375 4.988-2.625T13 12q0-3.075-2.013-5.325T6 4.05q-.425-.05-.713-.338T5 3q0-.425.3-.713T6 2.05q3.575.35 6.088 2.863T14.95 11h4.2l-.875-.9Q18 9.8 18 9.4t.3-.7q.275-.275.7-.275t.7.275l2.6 2.6q.15.15.213.325t.062.375q0 .2-.062.375t-.213.325l-2.6 2.6q-.275.275-.7.275t-.7-.275q-.3-.3-.3-.7t.275-.7l.875-.9h-4.2q-.35 3.575-2.862 6.088T6 21.95Z"/>`),
		g.Group(children),
	)
}