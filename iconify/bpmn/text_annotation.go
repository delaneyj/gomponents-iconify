package bpmn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextAnnotation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<g fill="none" stroke="currentColor"><path stroke-linecap="square" stroke-width="100.849" d="M1894.928-831.939h-458.41V852.603h458.41" transform="matrix(1.00396 0 0 1.01551 -3.915 995.337)"/><path stroke-dasharray="196.992 196.992" stroke-width="196.992" d="m201.143 840.592l1205.186-954.18" transform="matrix(1.00396 0 0 1.01551 -3.915 995.337)"/></g>`),
		g.Group(children),
	)
}