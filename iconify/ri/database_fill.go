package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DatabaseFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 7V4a1 1 0 0 1 1-1h9a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1h8Zm-6 9v2h5v-2H5Zm9 0v2h5v-2h-5Zm0-3v2h5v-2h-5Zm0-3v2h5v-2h-5Zm-9 3v2h5v-2H5Z"/>`),
		g.Group(children),
	)
}