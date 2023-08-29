package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TheaterComedy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 22q-2.5 0-4.25-1.75T1 16V9h12v7q0 2.5-1.75 4.25T7 22Zm-2-7.5q.425 0 .713-.288T6 13.5q0-.425-.288-.713T5 12.5q-.425 0-.713.288T4 13.5q0 .425.288.713T5 14.5Zm2 3.4q.975 0 1.738-.513T9.5 16h-5q0 .875.763 1.388T7 17.9Zm2-3.4q.425 0 .713-.288T10 13.5q0-.425-.288-.713T9 12.5q-.425 0-.713.288T8 13.5q0 .425.288.713T9 14.5Zm8 .5q-.65 0-1.288-.138T14.5 14.45V7.5H11V2h12v7q0 2.5-1.75 4.25T17 15Zm-2-7.5q.425 0 .713-.288T16 6.5q0-.425-.288-.713T15 5.5q-.425 0-.713.288T14 6.5q0 .425.288.713T15 7.5Zm-.5 3.4h5q0-.875-.763-1.388T17 9q-.85 0-1.675.45T14.5 10.9ZM19 7.5q.425 0 .713-.288T20 6.5q0-.425-.288-.713T19 5.5q-.425 0-.713.288T18 6.5q0 .425.288.713T19 7.5Z"/>`),
		g.Group(children),
	)
}