package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ranking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M17 18H4V42H17V18Z"/><path d="M30 6H17V42H30V6Z"/><path stroke-linecap="round" d="M43 26H30V42H43V26Z"/></g>`),
		g.Group(children),
	)
}