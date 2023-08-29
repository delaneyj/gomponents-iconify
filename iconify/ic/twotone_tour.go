package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneTour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 12V6h11.05l-1.2 3l1.2 3z" opacity=".3"/><path fill="currentColor" d="M21 4H7V2H5v20h2v-8h14l-2-5l2-5zM7 12V6h11.05l-1.2 3l1.2 3H7zm7-3c0 1.1-.9 2-2 2s-2-.9-2-2s.9-2 2-2s2 .9 2 2z"/>`),
		g.Group(children),
	)
}