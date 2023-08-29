package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveLocationOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.25 21l-1.4-1.4l1.575-1.6H15.25v-2h4.175l-1.575-1.6l1.4-1.4l4 4l-4 4ZM10 15q1.4 0 2.525-.688T14.3 12.5q-.875-.725-1.975-1.113T10 11q-1.225 0-2.325.388T5.7 12.5q.65 1.125 1.775 1.813T10 15Zm0-5q.825 0 1.413-.588T12 8q0-.825-.588-1.413T10 6q-.825 0-1.413.588T8 8q0 .825.588 1.413T10 10Zm0 1.675ZM10 22q-4.025-3.425-6.012-6.362T2 10.2q0-3.75 2.413-5.975T10 2q3.175 0 5.588 2.225T18 10.2q0 .225-.013.463t-.062.487H15.9q.05-.25.075-.488T16 10.2q0-2.725-1.738-4.463T10 4Q7.475 4 5.737 5.738T4 10.2q0 1.775 1.475 4.063T10 19.35q.575-.5 1.063-1t.937-.975l.225.225l.488.488q.262.262.475.487l.212.225q-.725.775-1.575 1.575T10 22Z"/>`),
		g.Group(children),
	)
}