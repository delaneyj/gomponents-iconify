package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BallBasketball(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 12a9 9 0 1 0 18 0a9 9 0 1 0-18 0m2.65-6.35l12.7 12.7m-12.7 0l12.7-12.7"/><path d="M12 3a9 9 0 0 0 9 9M3 12a9 9 0 0 1 9 9"/></g>`),
		g.Group(children),
	)
}