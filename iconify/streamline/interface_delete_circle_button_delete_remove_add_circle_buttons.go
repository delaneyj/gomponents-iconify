package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceDeleteCircleButtonDeleteRemoveAddCircleButtons(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M9.12 4.88L4.88 9.12m0-4.24l4.24 4.24"/><circle cx="7" cy="7" r="6.5"/></g>`),
		g.Group(children),
	)
}