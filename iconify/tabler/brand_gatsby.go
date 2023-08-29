package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandGatsby(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3.296 14.297l6.407 6.407a9.018 9.018 0 0 1-6.325-6.116l-.082-.291zM16 13h5c-.41 3.603-3.007 6.59-6.386 7.614L3.386 9.385A9 9 0 0 1 19.046 6.4"/>`),
		g.Group(children),
	)
}