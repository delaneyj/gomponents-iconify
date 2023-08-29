package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Superhero(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.1 10L18 2L7 10h2l-7 8l11-8h-1.9zm-4.3 1H3.9l2.5-1.8l7.5-5.5L10 2L3 5v3c0 2 .5 3.9 1.5 5.6L6.8 11zm6.4-2H16l-2.4 1.8L6.5 16c1 .9 2.2 1.6 3.5 2c4.2-1.5 7.1-5.5 7-10V5l-.2-.1L13.2 9z"/>`),
		g.Group(children),
	)
}