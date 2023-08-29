package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForwardToInboxRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.175 20H16q-.425 0-.713-.288T15 19q0-.425.288-.713T16 18h3.175l-.9-.9q-.3-.3-.287-.7t.312-.7q.3-.275.7-.287t.7.287l2.6 2.6q.3.3.3.7t-.3.7l-2.6 2.6q-.275.275-.687.288T18.3 22.3q-.275-.275-.275-.7t.275-.7l.875-.9ZM4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v7.8q-.675-.4-1.438-.6T19 13q-2.5 0-4.25 1.75T13 19v1H4Zm8-9L5.3 6.8q-.425-.275-.862-.025T4 7.525q0 .225.1.413t.3.312l7.075 4.425q.25.15.525.15t.525-.15L19.6 8.25q.2-.125.3-.313t.1-.412q0-.5-.437-.75T18.7 6.8L12 11Z"/>`),
		g.Group(children),
	)
}