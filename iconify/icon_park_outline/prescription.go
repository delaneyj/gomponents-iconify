package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Prescription(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linejoin="round" d="M8 8a2 2 0 0 1 2-2h17v12h13v22a2 2 0 0 1-2 2H10a2 2 0 0 1-2-2V8Z" clip-rule="evenodd"/><path stroke-linecap="round" stroke-linejoin="round" d="m27 6l13 12M27.024 6v12.082H40"/><path stroke-linecap="round" d="M14 30h12m-6-6v12"/></g>`),
		g.Group(children),
	)
}