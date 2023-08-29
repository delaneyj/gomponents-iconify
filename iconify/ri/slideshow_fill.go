package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlideshowFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 21v2h-2v-2H3a1 1 0 0 1-1-1V6h20v14a1 1 0 0 1-1 1h-8ZM8 10a3 3 0 1 0 3 3H8v-3Zm5 0v2h6v-2h-6Zm0 4v2h6v-2h-6ZM2 3h20v2H2V3Z"/>`),
		g.Group(children),
	)
}