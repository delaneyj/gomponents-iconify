package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageFlipHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M19 3v14h-8v3H9v-3H1V3h8V0h2v3h8zm-8.5 14V3h-1v14h1zM7 6.5L3 10l4 3.5v-7zM17 10l-4-3.5v7z"/>`),
		g.Group(children),
	)
}