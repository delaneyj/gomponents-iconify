package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aldi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.633 40.377h0c1.19 0 2.255-.738 2.673-1.853L30.197 9.475a2.855 2.855 0 0 1 2.673-1.852h0M10.077 40.377h0c1.19 0 2.255-.738 2.673-1.853L23.641 9.475a2.855 2.855 0 0 1 2.673-1.852h0M5.5 40.377L17.086 9.475a2.855 2.855 0 0 1 2.673-1.852h0m16.637 9.345a1.715 1.715 0 0 0-1.621-1.157h-3.443M28.262 24h9.546c.733 0 1.384.465 1.622 1.157m-14.238 7.032h15.687c.732 0 1.383.464 1.621 1.156"/>`),
		g.Group(children),
	)
}