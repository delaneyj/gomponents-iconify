package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WelcomeComments(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 2h10c1.1 0 2 .9 2 2v8c0 1.1-.9 2-2 2h-2l-5 5v-5H5c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2zm8.5 8.5L11 8l2.5-2.5l-1-1L10 7L7.5 4.5l-1 1L9 8l-2.5 2.5l1 1L10 9l2.5 2.5z"/>`),
		g.Group(children),
	)
}