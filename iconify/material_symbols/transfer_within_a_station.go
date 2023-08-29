package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransferWithinAStation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 23L5.8 8.9L4 9.6V13H2V8.3l5.05-2.15q.725-.3 1.475-.063T9.7 7l1 1.6q.675 1.075 1.788 1.738T15 11v2q-1.65 0-3.088-.688T9.5 10.5l-.6 3l2.1 2V23H9v-6l-2.1-2l-1.8 8H3ZM9.5 5.5q-.825 0-1.413-.588T7.5 3.5q0-.825.588-1.413T9.5 1.5q.825 0 1.413.588T11.5 3.5q0 .825-.588 1.413T9.5 5.5Zm10 17.5l-1.05-1.05l.7-.7H14v-1.5h5.15l-.7-.7L19.5 18l2.5 2.5l-2.5 2.5Zm-3-4.25l-2.5-2.5l2.5-2.5l1.05 1.05l-.7.7H22V17h-5.15l.7.7l-1.05 1.05Z"/>`),
		g.Group(children),
	)
}