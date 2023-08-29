package wi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoonWaxingGibbousThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 30 30"),
		g.Raw(`<path fill="currentColor" d="M11.8 14.43c0 2.39.24 4.52.71 6.39s1.31 3.5 2.51 4.89c1.52 0 2.98-.3 4.37-.89s2.59-1.4 3.6-2.4s1.81-2.2 2.4-3.6s.89-2.85.89-4.39s-.3-2.99-.89-4.38s-1.4-2.58-2.4-3.59s-2.2-1.81-3.6-2.4s-2.85-.89-4.37-.89C14 4.63 13.21 6.33 12.65 8.3s-.85 4-.85 6.13z"/>`),
		g.Group(children),
	)
}