package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Layout(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feLayout0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feLayout1" fill="currentColor" fill-rule="nonzero"><path id="feLayout2" d="M4 4h16a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2Zm0 2v5h4V6H4Zm0 7v5h4v-5H4Zm6-7v12h10V6H10Z"/></g></g>`),
		g.Group(children),
	)
}