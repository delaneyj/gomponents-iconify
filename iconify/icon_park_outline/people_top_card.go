package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PeopleTopCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M44 8H4v30h15l5 5l5-5h15V8Z"/><circle cx="24" cy="19" r="5"/><path d="M33 32c0-4.418-4.03-8-9-8s-9 3.582-9 8"/></g>`),
		g.Group(children),
	)
}