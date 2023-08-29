package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 15q-1.25 0-2.125-.875T1 12q0-1.25.875-2.125T4 9q.975 0 1.738.563T6.8 11H11v2H6.8q-.3.875-1.062 1.438T4 15Zm0-2q.425 0 .713-.288T5 12q0-.425-.288-.713T4 11q-.425 0-.713.288T3 12q0 .425.288.713T4 13Zm1 9v-2q3.325 0 5.663-2.337T13 12q0-3.325-2.337-5.663T5 4V2q3.925 0 6.75 2.6t3.2 6.4h4.2L17.6 9.4L19 8l4 4l-4 4l-1.4-1.4l1.55-1.6h-4.2q-.375 3.8-3.2 6.4T5 22Z"/>`),
		g.Group(children),
	)
}