package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Check(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feCheck0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feCheck1" fill="currentColor"><path id="feCheck2" d="m6 10l-2 2l6 6L20 8l-2-2l-8 8z"/></g></g>`),
		g.Group(children),
	)
}