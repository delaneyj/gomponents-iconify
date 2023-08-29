package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsInSimpleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M192 112h-48V64Zm-80 80v-48H64Z" opacity=".2"/><path d="M213.66 42.34a8 8 0 0 0-11.32 0L168 76.69l-18.34-18.35A8 8 0 0 0 136 64v48a8 8 0 0 0 8 8h48a8 8 0 0 0 5.66-13.66L179.31 88l34.35-34.34a8 8 0 0 0 0-11.32ZM152 104V83.31L172.69 104Zm-40 32H64a8 8 0 0 0-5.66 13.66L76.69 168l-34.35 34.34a8 8 0 0 0 11.32 11.32L88 179.31l18.34 18.35A8 8 0 0 0 120 192v-48a8 8 0 0 0-8-8Zm-8 36.69L83.31 152H104Z"/></g>`),
		g.Group(children),
	)
}