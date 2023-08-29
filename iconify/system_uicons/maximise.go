package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Maximise(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M18.5 16.5v-12a2 2 0 0 0-2-2h-12a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2z"/><path d="M2.5 14.5h10a2 2 0 0 0 2-2v-10"/><path d="M2.5 10.5h7a1 1 0 0 0 1-1v-7"/></g>`),
		g.Group(children),
	)
}