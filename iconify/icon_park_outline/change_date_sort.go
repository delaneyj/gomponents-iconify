package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChangeDateSort(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M6 5v25.004h36V5"/><path d="M15 23h4.002l13.85-14.143L28.993 5L15 19.143V23Z"/><path stroke-linecap="round" d="m30 37l-6 6l-6-6m6-7v13"/></g>`),
		g.Group(children),
	)
}