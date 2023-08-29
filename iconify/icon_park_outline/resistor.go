package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Resistor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M9.858 31.071a2 2 0 0 1 0-2.829l2.95-2.95a2 2 0 0 1 1.414-.585h2.585a2 2 0 0 0 1.415-.586l5.899-5.9a2 2 0 0 0 .586-1.413v-2.586a2 2 0 0 1 .586-1.415l2.95-2.95a2 2 0 0 1 2.828 0l7.07 7.072a2 2 0 0 1 0 2.828l-2.949 2.95a2 2 0 0 1-1.414.586h-2.586a2 2 0 0 0-1.414.586l-5.9 5.9a2 2 0 0 0-.585 1.413v2.586a2 2 0 0 1-.586 1.414l-2.95 2.95a2 2 0 0 1-2.828 0l-7.071-7.071Z"/><path stroke-linecap="round" d="m7.03 40.97l6.363-6.364m21.214-21.213l6.364-6.364"/></g>`),
		g.Group(children),
	)
}