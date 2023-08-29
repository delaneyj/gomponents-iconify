package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BusinessChipOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 16h6q.425 0 .713-.288T16 15v-4q0-.425-.288-.713T15 10h-1V9q0-.425-.288-.713T13 8h-2q-.425 0-.713.288T10 9v1H9q-.425 0-.713.288T8 11v4q0 .425.288.713T9 16Zm2-6V9h2v1h-2Zm-3 9q-2.925 0-4.963-2.038T1 12q0-2.925 2.038-4.963T8 5h8q2.925 0 4.963 2.038T23 12q0 2.925-2.038 4.963T16 19H8Zm0-2h8q2.075 0 3.538-1.463T21 12q0-2.075-1.463-3.538T16 7H8Q5.925 7 4.462 8.463T3 12q0 2.075 1.463 3.538T8 17Zm4-5Z"/>`),
		g.Group(children),
	)
}