package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 10v4h6v-4h-6Zm-2 0H5v4h6v-4Zm2 9h6v-3h-6v3Zm-2 0v-3H5v3h6Zm2-14v3h6V5h-6Zm-2 0H5v3h6V5ZM4 3h16a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1Z"/>`),
		g.Group(children),
	)
}