package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardBoxFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 3h18a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1Zm2 4v2h2V7H5Zm0 4v2h2v-2H5Zm0 4v2h14v-2H5Zm4-4v2h2v-2H9Zm0-4v2h2V7H9Zm4 0v2h2V7h-2Zm4 0v2h2V7h-2Zm-4 4v2h2v-2h-2Zm4 0v2h2v-2h-2Z"/>`),
		g.Group(children),
	)
}