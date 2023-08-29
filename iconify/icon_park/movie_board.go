package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MovieBoard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M44 16H4V42H44V16Z"/><path stroke="#000" d="M44 16V6H4V16H44Z"/><path stroke="#000" d="M26 6L30 16"/><path stroke="#000" d="M18 6L22 16"/><path stroke="#000" d="M10 6L14 16"/><path stroke="#000" d="M34 6L38 16"/><path stroke="#fff" d="M12 24H36"/><path stroke="#fff" d="M12 32H24"/></g>`),
		g.Group(children),
	)
}