package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BikeScooterRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 19v-2h4.1q.275-1.7 1.413-2.975T8.3 12.25L6.85 6H4q-.425 0-.713-.288T3 5q0-.425.288-.713T4 4h2.85q.675 0 1.238.438T8.8 5.55L10.75 14H10q-1.65 0-2.825 1.175T6 18v1H0Zm10 2q-1.25 0-2.125-.875T7 18q0-1.25.875-2.125T10 15q1.25 0 2.125.875T13 18q0 1.25-.875 2.125T10 21Zm9-3q-1.8 0-3.163-1.113T14.1 14h-2.35l-.45-2h2.8q.125-.575.338-1.075T15 10h-4.15l-.45-2h5.65l-1.1-3h-1.6q-.425 0-.712-.288T12.35 4q0-.425.288-.713T13.35 3h1.6q.65 0 1.163.35t.737.95l1.35 3.65h.8q2.075 0 3.538 1.463T24 12.95q0 2.125-1.463 3.588T19 18Zm-1.95-7.3l.6 1.7q.125.4.513.575t.787.025q.4-.15.575-.525t.025-.775L18.9 10l-1.85.7Z"/>`),
		g.Group(children),
	)
}