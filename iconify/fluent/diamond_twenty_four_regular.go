package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiamondTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.659 13.593a2.25 2.25 0 0 1 0-3.182l7.752-7.752a2.25 2.25 0 0 1 3.182 0l7.751 7.752a2.25 2.25 0 0 1 0 3.182l-7.751 7.751a2.25 2.25 0 0 1-3.182 0l-7.752-7.751Zm1.06-2.122a.75.75 0 0 0 0 1.06l7.752 7.753a.75.75 0 0 0 1.06 0l7.753-7.752a.75.75 0 0 0 0-1.06L12.532 3.72a.75.75 0 0 0-1.06 0L3.72 11.47Z"/>`),
		g.Group(children),
	)
}