package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBendDoubleUpRightDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="m176 104l-48 48V56Z" opacity=".2"/><path d="m229.66 109.66l-48 48a8 8 0 0 1-11.32-11.32L212.69 104l-42.35-42.34a8 8 0 0 1 11.32-11.32l48 48a8 8 0 0 1 0 11.32Zm-48 0l-48 48A8 8 0 0 1 120 152v-39.63A88.11 88.11 0 0 0 40 200a8 8 0 0 1-16 0a104.15 104.15 0 0 1 96-103.7V56a8 8 0 0 1 13.66-5.66l48 48a8 8 0 0 1 0 11.32Zm-17-5.66L136 75.31v57.38Z"/></g>`),
		g.Group(children),
	)
}