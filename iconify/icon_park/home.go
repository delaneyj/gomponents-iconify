package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Home(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><path fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M9 18V42H39V18L24 6L9 18Z"/><path fill="#43CCF8" stroke="#fff" stroke-linejoin="round" d="M19 29V42H29V29H19Z"/><path stroke="#000" stroke-linecap="round" d="M9 42H39"/></g>`),
		g.Group(children),
	)
}