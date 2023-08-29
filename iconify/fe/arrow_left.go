package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feArrowLeft0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feArrowLeft1" fill="currentColor"><path id="feArrowLeft2" d="m15 4l2 2l-6 6l6 6l-2 2l-8-8z"/></g></g>`),
		g.Group(children),
	)
}