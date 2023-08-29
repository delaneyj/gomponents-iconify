package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleLarge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Z"/><path fill-rule="evenodd" d="M24 41.5c9.665 0 17.5-7.835 17.5-17.5S33.665 6.5 24 6.5S6.5 14.335 6.5 24S14.335 41.5 24 41.5Zm0 2.5c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}