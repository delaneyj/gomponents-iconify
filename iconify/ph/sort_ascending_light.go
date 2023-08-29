package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortAscendingLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M126 128a6 6 0 0 1-6 6H48a6 6 0 0 1 0-12h72a6 6 0 0 1 6 6ZM48 70h136a6 6 0 0 0 0-12H48a6 6 0 0 0 0 12Zm56 116H48a6 6 0 0 0 0 12h56a6 6 0 0 0 0-12Zm124.24-22.24a6 6 0 0 0-8.48 0L190 193.51V112a6 6 0 0 0-12 0v81.51l-29.76-29.75a6 6 0 0 0-8.48 8.48l40 40a6 6 0 0 0 8.48 0l40-40a6 6 0 0 0 0-8.48Z"/>`),
		g.Group(children),
	)
}