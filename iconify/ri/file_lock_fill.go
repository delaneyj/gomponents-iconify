package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileLockFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16 2l5 5v14.008a.993.993 0 0 1-.993.992H3.993A1 1 0 0 1 3 21.008V2.992C3 2.444 3.445 2 3.993 2H16Zm-1 9v-1a3 3 0 1 0-6 0v1H8v5h8v-5h-1Zm-2 0h-2v-1a1 1 0 1 1 2 0v1Z"/>`),
		g.Group(children),
	)
}