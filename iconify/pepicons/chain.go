package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chain(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><rect width="6" height="10" x="12.784" y="2.384" rx="3" transform="rotate(33.038 12.784 2.384)"/><rect width="6" height="10" x="7.836" y="6.323" rx="3" transform="rotate(33.038 7.836 6.323)"/></g>`),
		g.Group(children),
	)
}