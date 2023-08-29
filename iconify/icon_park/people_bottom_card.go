package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PeopleBottomCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M44 39H4V9H19L24 4L29 9H44V39Z"/><circle cx="24" cy="20" r="5" fill="#2F88FF"/><path d="M33 33C33 28.5817 28.9706 25 24 25C19.0294 25 15 28.5817 15 33"/></g>`),
		g.Group(children),
	)
}