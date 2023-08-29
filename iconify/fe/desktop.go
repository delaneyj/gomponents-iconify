package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Desktop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feDesktop0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feDesktop1" fill="currentColor"><path id="feDesktop2" d="M13 18h2a1 1 0 0 1 0 2H9a1 1 0 0 1 0-2h2v-2H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2h-7v2ZM4 6v8h16V6H4Z"/></g></g>`),
		g.Group(children),
	)
}