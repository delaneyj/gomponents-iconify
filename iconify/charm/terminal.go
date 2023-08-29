package charm

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Terminal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 2.75h12.5v10.5H1.75z"/><path d="M8.75 10.25h2.5m-6.5-4.5L7.25 8l-2.5 2.25"/></g>`),
		g.Group(children),
	)
}