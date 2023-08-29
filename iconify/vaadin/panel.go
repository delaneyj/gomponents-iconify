package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Panel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 0v16h16V0H0zm13 15H1V3h12v12zm2 0h-1v-1h1v1zm0-2h-1V5h1v8zm0-9h-1V3h1v1z"/>`),
		g.Group(children),
	)
}