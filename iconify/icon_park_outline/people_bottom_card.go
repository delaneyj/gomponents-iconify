package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PeopleBottomCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M44 39H4V9h15l5-5l5 5h15v30Z"/><circle cx="24" cy="20" r="5"/><path d="M33 33c0-4.418-4.03-8-9-8s-9 3.582-9 8"/></g>`),
		g.Group(children),
	)
}