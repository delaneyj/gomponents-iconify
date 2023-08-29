package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TokenOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21.7q-.25 0-.5-.062t-.475-.188l-7-3.875q-.475-.275-.75-.738T3 15.825v-7.65q0-.55.275-1.012t.75-.738l7-3.875q.225-.125.475-.187T12 2.3q.25 0 .5.063t.475.187l7 3.875q.475.275.75.738T21 8.175v7.65q0 .55-.275 1.013t-.75.737l-7 3.875q-.225.125-.475.188t-.5.062ZM9.1 9.25q.55-.575 1.288-.913T12 8q.875 0 1.613.338t1.287.912l3-1.675L12 4.3L6.1 7.575l3 1.675Zm1.9 9.9v-3.275q-1.3-.35-2.15-1.413T8 12q0-.275.025-.525t.1-.475L5 9.25v6.575l6 3.325ZM12 14q.825 0 1.413-.588T14 12q0-.825-.588-1.413T12 10q-.825 0-1.413.588T10 12q0 .825.588 1.413T12 14Zm1 5.15l6-3.325V9.25L15.875 11q.075.25.1.488T16 12q0 1.4-.85 2.463T13 15.875v3.275Z"/>`),
		g.Group(children),
	)
}