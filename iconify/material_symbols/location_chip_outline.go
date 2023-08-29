package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationChipOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 16l.513-.438q.512-.437 1.112-1.124t1.113-1.526q.512-.837.512-1.662q0-1.35-.95-2.3T12 8q-1.35 0-2.3.95t-.95 2.3q0 .825.513 1.663t1.112 1.524q.6.688 1.113 1.126L12 16Zm0-4q-.425 0-.713-.288T11 11q0-.425.288-.713T12 10q.425 0 .713.288T13 11q0 .425-.288.713T12 12Zm-4 7q-2.925 0-4.963-2.038T1 12q0-2.925 2.038-4.963T8 5h8q2.925 0 4.963 2.038T23 12q0 2.925-2.038 4.963T16 19H8Zm0-2h8q2.075 0 3.538-1.463T21 12q0-2.075-1.463-3.538T16 7H8Q5.925 7 4.462 8.463T3 12q0 2.075 1.463 3.538T8 17Zm4-5Z"/>`),
		g.Group(children),
	)
}