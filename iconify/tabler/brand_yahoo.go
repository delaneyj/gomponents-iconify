package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandYahoo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 6h5M7 18h7M4.5 6l5.5 7v5m0-5l6-5m-3.5 0h5m2.5 3v4m0 3v.01"/>`),
		g.Group(children),
	)
}