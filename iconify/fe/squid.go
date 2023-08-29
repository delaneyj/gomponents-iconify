package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Squid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feSquid0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feSquid1" fill="currentColor"><path id="feSquid2" d="M8 10L6 8l6-6l6 6l-2 2v7a1 1 0 0 0 2 0a1 1 0 0 1 2 0a3 3 0 0 1-5 2.236V21a1 1 0 0 1-2 0v-2h-2v2a1 1 0 0 1-2 0v-1.764A3 3 0 0 1 4 17a1 1 0 0 1 2 0a1 1 0 0 0 2 0v-7Z"/></g></g>`),
		g.Group(children),
	)
}