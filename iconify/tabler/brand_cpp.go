package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandCpp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 12h4m-2-2v4m-9-2h4m-2-2v4M9 9a3 3 0 0 0-3-3h-.5A3.5 3.5 0 0 0 2 9.5v5A3.5 3.5 0 0 0 5.5 18H6a3 3 0 0 0 3-3"/>`),
		g.Group(children),
	)
}