package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RectangularVertebra(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M12 9L24 4L36 9L44 38L24 44L4 38L12 9Z"/><path d="M24 44L24 14"/><path d="M36 9L24 14"/><path d="M12 9L24 14"/><path d="M5 38L24 32"/><path d="M44 38L24 32"/></g>`),
		g.Group(children),
	)
}