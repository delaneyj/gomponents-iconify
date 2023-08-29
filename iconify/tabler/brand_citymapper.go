package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandCitymapper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 11a1 1 0 1 1-1 1.013a1 1 0 0 1 1-1V11zm18 0a1 1 0 1 1-1 1.013a1 1 0 0 1 1-1V11zM8 12h8m-3-3l3 3l-3 3"/>`),
		g.Group(children),
	)
}