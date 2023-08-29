package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileMovie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M10 0H2v16h12V4l-4-4zM9 5h4v10H3V1h6v4zm1-1V1l3 3h-3z"/><path fill="currentColor" d="M10 10V8H4v5h6v-2l2 2V8z"/>`),
		g.Group(children),
	)
}