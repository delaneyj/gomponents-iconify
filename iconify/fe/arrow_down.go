package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feArrowDown0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feArrowDown1" fill="currentColor"><path id="feArrowDown2" d="m6 7l6 6l6-6l2 2l-8 8l-8-8z"/></g></g>`),
		g.Group(children),
	)
}