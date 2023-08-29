package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mascara(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-width="4"><rect width="12" height="28" x="28" y="16" stroke="#000" stroke-linejoin="round"/><rect width="12" height="16" x="8" y="28" fill="#2F88FF" stroke="#000" stroke-linejoin="round"/><path stroke="#000" stroke-linejoin="round" d="M14 4V28"/><path stroke="#fff" d="M20 36H8"/><path stroke="#000" stroke-linejoin="round" d="M20 32V40"/><path stroke="#000" stroke-linejoin="round" d="M8 32V40"/><path stroke="#000" stroke-linejoin="round" d="M18 9H10"/><path stroke="#000" stroke-linejoin="round" d="M20 15L8 15"/><path stroke="#000" stroke-linejoin="round" d="M18 21H10"/></g>`),
		g.Group(children),
	)
}