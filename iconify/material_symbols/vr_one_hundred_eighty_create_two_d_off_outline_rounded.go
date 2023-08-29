package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VrOneHundredEightyCreateTwoDOffOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22 19.175l-2-2V12h-5.175l-2-2H20q.825 0 1.413.588T22 12v7.175ZM14.6 17.4ZM12 22q-.825 0-1.413-.588T10 20v-7.2L4.6 7.4q-.275.6-.437 1.25T4 10q0 1.925 1.1 3.463T8 15.65v2.1q-2.675-.675-4.338-2.838T2 10q0-1.125.3-2.15t.825-1.925l-1.8-1.8q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275L21.2 21.2q.3.3.275.725t-.325.725q-.275.275-.637.288t-.638-.263L19.2 22H12ZM9.175 9.175ZM13 19q-.15 0-.225-.138t.025-.262l1.3-1.725q.075-.1.2-.1t.2.1l1 1.375l1.175-1.575L19 19h-6Zm-1 1h5.2L12 14.8V20Zm5.425-5.425ZM15.65 8Q15 6.2 13.462 5.1T10 4q-.675 0-1.325.15t-1.25.425L5.95 3.1q.925-.55 1.95-.825T10 2q2.75 0 4.913 1.663T17.75 8h-2.1Zm-4.675.125Z"/>`),
		g.Group(children),
	)
}