package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandPowershell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4.887 20h11.868c.893 0 1.664-.665 1.847-1.592l2.358-12c.212-1.081-.442-2.14-1.462-2.366A1.784 1.784 0 0 0 19.113 4H7.245c-.893 0-1.664.665-1.847 1.592l-2.358 12c-.212 1.081.442 2.14 1.462 2.366c.127.028.256.042.385.042z"/><path d="m9 8l4 4l-6 4m5 0h3"/></g>`),
		g.Group(children),
	)
}