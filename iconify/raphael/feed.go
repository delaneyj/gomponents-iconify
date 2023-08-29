package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Feed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4.135 16.762c3.078 0 5.972 1.205 8.146 3.39a11.543 11.543 0 0 1 3.378 8.203h4.745c0-9.008-7.3-16.335-16.27-16.335v4.742zm.006-8.408c10.974 0 19.9 8.975 19.9 20.006h4.742c0-13.646-11.054-24.75-24.642-24.75v4.744zm6.56 16.69a3.286 3.286 0 1 1-6.572.002a3.286 3.286 0 0 1 6.572-.001z"/>`),
		g.Group(children),
	)
}