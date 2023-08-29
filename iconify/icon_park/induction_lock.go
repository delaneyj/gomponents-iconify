package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InductionLock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M9 18V42H39V18L24 6L9 18Z"/><path stroke="#fff" d="M24 22V34"/><path stroke="#fff" d="M31 26V30"/><path stroke="#fff" d="M17 26V30"/></g>`),
		g.Group(children),
	)
}