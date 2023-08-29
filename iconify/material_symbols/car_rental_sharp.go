package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarRentalSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 7q-1.25 0-2.125-.875T5 4q0-1.25.875-2.125T8 1q.95 0 1.725.563T10.85 3H19v2h-1v2h-2V5h-5.15q-.35.875-1.125 1.438T8 7Zm0-2q.425 0 .713-.288T9 4q0-.425-.288-.713T8 3q-.425 0-.713.288T7 4q0 .425.288.713T8 5Zm1 12.5q.425 0 .713-.288T10 16.5q0-.425-.288-.713T9 15.5q-.425 0-.713.288T8 16.5q0 .425.288.713T9 17.5Zm6 0q.425 0 .713-.288T16 16.5q0-.425-.288-.713T15 15.5q-.425 0-.713.288T14 16.5q0 .425.288.713T15 17.5ZM5 22v-7.4L6.925 9h10.15L19 14.6V22h-2v-2H7v2H5Zm2.65-9h8.7l-.7-2h-7.3l-.7 2Z"/>`),
		g.Group(children),
	)
}