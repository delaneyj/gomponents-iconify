package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VtbTwentyFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.907 20.549H6.592M42.5 8.599H15.736C11.087 12.825 7.284 16.769 5.5 24c1.784 7.231 5.588 11.175 10.236 15.401H42.5c0-5.588-8.593-13.523-8.593-13.523H11.604m27.515-11.304H9.994"/>`),
		g.Group(children),
	)
}