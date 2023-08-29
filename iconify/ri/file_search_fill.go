package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileSearchFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16 2l5 5v14.008a.993.993 0 0 1-.993.992H3.993A1 1 0 0 1 3 21.008V2.992C3 2.444 3.445 2 3.993 2H16Zm-2.471 12.446l2.21 2.21l1.415-1.413l-2.21-2.21a4.001 4.001 0 0 0-6.276-4.861a4 4 0 0 0 4.861 6.274Zm-.618-2.032a2 2 0 1 1-2.828-2.828a2 2 0 0 1 2.828 2.828Z"/>`),
		g.Group(children),
	)
}