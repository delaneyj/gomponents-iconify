package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShapesOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 15Zm-7 2.95q.25.025.488.038T9 18q.275 0 .513-.013T10 17.95V20h10V10h-2.05q.025-.25.038-.488T18 9q0-.275-.013-.513T17.95 8H20q.825 0 1.413.588T22 10v10q0 .825-.588 1.413T20 22H10q-.825 0-1.413-.588T8 20v-2.05ZM9 16q-2.925 0-4.963-2.038T2 9q0-2.925 2.038-4.963T9 2q2.925 0 4.963 2.038T16 9q0 2.925-2.038 4.963T9 16Zm0-2q2.075 0 3.538-1.463T14 9q0-2.075-1.463-3.538T9 4Q6.925 4 5.462 5.463T4 9q0 2.075 1.463 3.538T9 14Zm0-5Z"/>`),
		g.Group(children),
	)
}