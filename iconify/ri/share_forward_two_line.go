package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareForwardTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 19h16v-5h2v6a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1v-6h2v5ZM16.172 7l-3.95-3.95l1.414-1.414L20 8l-6.364 6.364l-1.414-1.415L16.172 9H5V7h11.172Z"/>`),
		g.Group(children),
	)
}