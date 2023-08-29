package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Safe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M1 0v16h3v-1h8v1h3V0H1zm13 10h-1V5h1v5zm0-7h-1V2H3v11h10v-1h1v2H2V1h12v2zM8.5 7.5c0 1.1-.9 2-2 2s-2-.9-2-2s.9-2 2-2s2 .9 2 2z"/><path fill="currentColor" d="M7.5 7.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0z"/>`),
		g.Group(children),
	)
}