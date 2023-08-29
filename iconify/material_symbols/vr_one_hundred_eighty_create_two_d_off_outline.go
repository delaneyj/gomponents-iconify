package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VrOneHundredEightyCreateTwoDOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22 19.175l-2-2V12h-5.175l-2-2H20q.825 0 1.413.588T22 12v7.175ZM15.65 8Q15 6.2 13.462 5.1T10 4q-.675 0-1.325.15t-1.25.425L5.95 3.1q.925-.55 1.95-.825T10 2q2.75 0 4.913 1.663T17.75 8h-2.1Zm-1.05 9.4Zm5.875 5.9L19.2 22H12q-.825 0-1.413-.588T10 20v-7.2L4.6 7.4q-.275.6-.437 1.25T4 10q0 1.925 1.1 3.463T8 15.65v2.1q-2.675-.675-4.338-2.838T2 10q0-1.125.3-2.15t.825-1.925l-2.5-2.5l1.4-1.4L21.9 21.9l-1.425 1.4Zm-9.5-15.175Zm-1.8 1.05ZM12.5 19l1.8-2.4l1.2 1.65l1.175-1.575L19 19h-6.5Zm-.5 1h5.2L12 14.8V20Zm5.425-5.425Z"/>`),
		g.Group(children),
	)
}