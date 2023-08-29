package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Duplicate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M18.5 12.5v-8a2 2 0 0 0-2-2h-8a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2z"/><path d="M6.5 6.503h-2a2 2 0 0 0-2 2V16.5a2 2 0 0 0 2 2h.003l8-.014a2 2 0 0 0 1.997-2v-1.983m-2-9.003v6m3-3h-6"/></g>`),
		g.Group(children),
	)
}