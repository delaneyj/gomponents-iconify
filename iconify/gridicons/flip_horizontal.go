package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 18v-5h3v-2h-3V6a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v5H1v2h3v5a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2zM6 6h12v5H6V6z"/>`),
		g.Group(children),
	)
}