package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feArrowRight0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feArrowRight1" fill="currentColor"><path id="feArrowRight2" d="m9.005 4l8 8l-8 8L7 18l6.005-6L7 6z"/></g></g>`),
		g.Group(children),
	)
}