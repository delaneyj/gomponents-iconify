package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Headphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feHeadphone0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feHeadphone1" fill="currentColor"><path id="feHeadphone2" d="M19 13a2 2 0 0 1 2 2v4a2 2 0 1 1-4 0v-9a5 5 0 0 0-10 0v9a2 2 0 1 1-4 0v-4a2 2 0 0 1 2-2v-3a7 7 0 1 1 14 0v3Z"/></g></g>`),
		g.Group(children),
	)
}