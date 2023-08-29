package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VipDiamondLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.878 3.003h14.254a1 1 0 0 1 .809.412l3.822 5.256a.5.5 0 0 1-.037.633l-11.354 12.3a.5.5 0 0 1-.706.03L.283 9.303a.5.5 0 0 1-.037-.633l3.823-5.256a1 1 0 0 1 .809-.412Zm.509 2l-2.8 3.85l9.418 10.201l9.417-10.202l-2.8-3.85H5.388Z"/>`),
		g.Group(children),
	)
}