package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipCameraAndroidRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-3.2 0-5.763-1.825T2.6 15.4q-.125-.375.075-.7t.6-.45q.4-.125.763.088t.512.587q.9 2.275 2.925 3.675T12 20q2.15 0 4-1.063T18.9 16H17q-.425 0-.712-.288T16 15q0-.425.288-.713T17 14h4q.425 0 .713.288T22 15v4q0 .425-.288.713T21 20q-.425 0-.713-.288T20 19v-1q-1.425 1.9-3.525 2.95T12 22Zm0-18Q9.85 4 8 5.063T5.1 8H7q.425 0 .713.288T8 9q0 .425-.288.713T7 10H3q-.425 0-.713-.288T2 9V5q0-.425.288-.713T3 4q.425 0 .713.288T4 5v1q1.425-1.9 3.525-2.95T12 2q3.2 0 5.763 1.825T21.4 8.6q.125.375-.075.7t-.6.45q-.4.125-.763-.088t-.512-.587Q18.55 6.8 16.525 5.4T12 4Zm0 11q-1.25 0-2.125-.875T9 12q0-1.25.875-2.125T12 9q1.25 0 2.125.875T15 12q0 1.25-.875 2.125T12 15Z"/>`),
		g.Group(children),
	)
}