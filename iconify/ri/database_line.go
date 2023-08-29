package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DatabaseLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 19V9H4v10h7Zm0-12V4a1 1 0 0 1 1-1h9a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1h8Zm2-2v14h7V5h-7ZM5 16h5v2H5v-2Zm9 0h5v2h-5v-2Zm0-3h5v2h-5v-2Zm0-3h5v2h-5v-2Zm-9 3h5v2H5v-2Z"/>`),
		g.Group(children),
	)
}