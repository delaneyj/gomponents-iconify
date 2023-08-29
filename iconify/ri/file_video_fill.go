package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileVideoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16 2l5 5v14.008a.993.993 0 0 1-.993.992H3.993A1 1 0 0 1 3 21.008V2.992C3 2.444 3.445 2 3.993 2H16Zm-1 9.667l-4.878-3.253a.4.4 0 0 0-.622.333v6.505a.4.4 0 0 0 .622.333L15 12.333a.401.401 0 0 0 0-.666Z"/>`),
		g.Group(children),
	)
}