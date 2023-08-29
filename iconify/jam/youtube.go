package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Youtube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.812.017H4.145C1.855.017 0 1.852 0 4.116v5.768c0 2.264 1.856 4.1 4.145 4.1h11.667c2.29 0 4.145-1.836 4.145-4.1V4.116c0-2.264-1.856-4.1-4.145-4.1zM13.009 7.28L7.552 9.855a.219.219 0 0 1-.314-.196V4.35a.22.22 0 0 1 .318-.193l5.458 2.735a.216.216 0 0 1-.005.389z"/>`),
		g.Group(children),
	)
}