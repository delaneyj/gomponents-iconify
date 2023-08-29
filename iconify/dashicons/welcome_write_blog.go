package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WelcomeWriteBlog(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m16.89 1.2l1.41 1.41c.39.39.39 1.02 0 1.41L14 8.33V18H3V3h10.67l1.8-1.8c.4-.39 1.03-.4 1.42 0zm-5.66 8.48l5.37-5.36l-1.42-1.42l-5.36 5.37l-.71 2.12z"/>`),
		g.Group(children),
	)
}