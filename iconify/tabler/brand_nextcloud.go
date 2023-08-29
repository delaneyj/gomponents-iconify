package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandNextcloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 12a5 5 0 1 0 10 0a5 5 0 1 0-10 0m-5 .5a2.5 2.5 0 1 0 5 0a2.5 2.5 0 1 0-5 0m15 0a2.5 2.5 0 1 0 5 0a2.5 2.5 0 1 0-5 0"/>`),
		g.Group(children),
	)
}