package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrownThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#2F88FF" stroke="#000" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M13 42H35L41 21L31 26L24 12L17 26L7 21L13 42Z"/><circle cx="7" cy="18" r="3"/><circle cx="24" cy="9" r="3"/><circle cx="41" cy="18" r="3"/></g>`),
		g.Group(children),
	)
}