package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EyeCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 12a2 2 0 1 0 4 0a2 2 0 0 0-4 0"/><path d="M11.11 17.958C7.901 17.651 5.2 15.665 3 12c2.4-4 5.4-6 9-6c3.6 0 6.6 2 9 6c-.21.352-.427.688-.647 1.008M20 21l2-2l-2-2m-3 0l-2 2l2 2"/></g>`),
		g.Group(children),
	)
}