package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TempleBuddhistOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22v-9.15q-1.3-.325-2.15-1.375T1 9.025h2q0 .8.588 1.388T4.974 11H6V8.85q-1.3-.325-2.15-1.375T3 5.025h2q0 .8.588 1.387T6.975 7H7.5L12 1l4.5 6h.525q.8 0 1.388-.588T19 5.025h2q0 1.4-.85 2.45T18 8.85V11h1.025q.8 0 1.388-.587T21 9.025h2q0 1.4-.85 2.45T20 12.85V22h-7v-4q0-.425-.288-.713T12 17q-.425 0-.713.288T11 18v4H4Zm6-15h4l-2-2.675L10 7Zm-2 4h8V9H8v2Zm-2 9h3v-2q0-1.25.875-2.125T12 15q1.25 0 2.125.875T15 18v2h3v-7H6v7Zm6-7Zm0-6Zm0 4Z"/>`),
		g.Group(children),
	)
}