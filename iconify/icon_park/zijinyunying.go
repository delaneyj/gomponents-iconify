package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zijinyunying(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M34 6H14L3 24L14 42H34L45 24L34 6Z"/><path fill="#43CCF8" stroke="#fff" d="M15 29L24 15L33 29H15Z"/></g>`),
		g.Group(children),
	)
}