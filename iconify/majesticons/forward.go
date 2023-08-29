package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M12 5a1 1 0 0 1 1.707-.707l7 7a1 1 0 0 1 0 1.414l-7 7A1 1 0 0 1 12 19v-3H4a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1h8V5z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}