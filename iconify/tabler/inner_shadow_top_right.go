package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InnerShadowTopRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18z"/><path d="M18 12a6 6 0 0 0-6-6"/></g>`),
		g.Group(children),
	)
}