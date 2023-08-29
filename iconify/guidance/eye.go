package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eye(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor"><path d="M6.106 7.053a8.079 8.079 0 0 1 11.788 0L20.5 10c1 1 2.224 2 3.5 2c-1.276 0-2.5 1-3.5 2l-2.606 2.947a8.079 8.079 0 0 1-11.788 0L3.5 14c-1-1-2.224-2-3.5-2c1.276 0 2.5-1 3.5-2l2.606-2.947Z"/><path d="M9.5 12a2.5 2.5 0 1 1 5 0a2.5 2.5 0 0 1-5 0Z"/></g>`),
		g.Group(children),
	)
}