package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBendRightDownDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="m200 176l-48 48l-48-48Z" opacity=".2"/><path d="M207.39 172.94A8 8 0 0 0 200 168h-40v-40A104.11 104.11 0 0 0 56 24a8 8 0 0 0 0 16a88.1 88.1 0 0 1 88 88v40h-40a8 8 0 0 0-5.66 13.66l48 48a8 8 0 0 0 11.32 0l48-48a8 8 0 0 0 1.73-8.72ZM152 212.69L123.31 184h57.38Z"/></g>`),
		g.Group(children),
	)
}