package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BallBaseball(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5.636 18.364A9 9 0 1 0 18.364 5.636A9 9 0 0 0 5.636 18.364z"/><path d="M12.495 3.02a9 9 0 0 1-9.475 9.475m17.96-.99a9 9 0 0 0-9.475 9.475M9 9l2 2m2 2l2 2m-4-8l2 1m-6 3l1 2m8-2l1 2m-6 3l2 1"/></g>`),
		g.Group(children),
	)
}