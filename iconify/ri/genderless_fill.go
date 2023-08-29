package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GenderlessFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 7.066V1h2v6.066A7.501 7.501 0 0 1 12 22a7.5 7.5 0 0 1-1-14.934Z"/>`),
		g.Group(children),
	)
}