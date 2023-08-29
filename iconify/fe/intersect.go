package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Intersect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feIntersect0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feIntersect1" fill="currentColor"><path id="feIntersect2" d="M15 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-8a2 2 0 0 1 2-2h4V5a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2h-4Zm-5-1h4v-4h-4v4Z"/></g></g>`),
		g.Group(children),
	)
}