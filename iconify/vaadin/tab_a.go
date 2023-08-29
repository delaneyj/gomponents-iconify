package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabA(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M9 10H0V6h9V4l5 4l-5 4v-2zm5-6h2v8h-2V4z"/>`),
		g.Group(children),
	)
}