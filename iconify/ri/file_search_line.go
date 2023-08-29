package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileSearchLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 4H5v16h14V8h-4V4ZM3 2.992C3 2.444 3.447 2 3.998 2H16l5 5v13.992A1 1 0 0 1 20.007 22H3.993A1 1 0 0 1 3 21.008V2.992Zm10.529 11.454a4.001 4.001 0 0 1-4.86-6.274a4 4 0 0 1 6.274 4.86l2.21 2.21l-1.413 1.415l-2.211-2.21Zm-.618-2.032a2 2 0 1 0-2.828-2.828a2 2 0 0 0 2.828 2.828Z"/>`),
		g.Group(children),
	)
}