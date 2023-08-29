package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UtdRooms(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m13.249 4.518l18.133 2.57l.007 36.412l-18.144-2.57Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m13.245 4.5l21.51.073l-.073 35.781l-3.22.074"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.594 26.018c.058 1.54-.644 2.547-1.375 2.265c-.594-.228-1.014-1.222-1.042-2.275c-.027-1.039.281-2.225.903-2.298c.839-.099 1.457.796 1.514 2.308Z"/>`),
		g.Group(children),
	)
}