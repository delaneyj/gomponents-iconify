package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AcceptEmail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M44 24V9H24H4V24V39H24"/><path d="M44 34L30 34"/><path d="M35 29L30 34L35 39"/><path d="M4 9L24 24L44 9"/></g>`),
		g.Group(children),
	)
}