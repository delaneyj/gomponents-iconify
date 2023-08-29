package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandGolang(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.695 14.305c1.061 1.06 2.953.888 4.226-.384c1.272-1.273 1.444-3.165.384-4.226c-1.061-1.06-2.953-.888-4.226.384c-1.272 1.273-1.444 3.165-.384 4.226zM12.68 9.233c-1.084-.497-2.545-.191-3.591.846c-1.284 1.273-1.457 3.165-.388 4.226c1.07 1.06 2.978.888 4.261-.384A3.669 3.669 0 0 0 14 12h-2.427M5.5 15H4m2-6H4m1 3H2"/>`),
		g.Group(children),
	)
}