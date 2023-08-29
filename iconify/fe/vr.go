package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feVr0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feVr1" fill="currentColor" fill-rule="nonzero"><path id="feVr2" d="m15 17l-2-2h-2l-2 2H7a2 2 0 0 1-2-2V9a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2h-2Zm7-3a1 1 0 0 1-2 0v-4a1 1 0 0 1 2 0v4ZM4 14a1 1 0 0 1-2 0v-4a1 1 0 1 1 2 0v4Z"/></g></g>`),
		g.Group(children),
	)
}