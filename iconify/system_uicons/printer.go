package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Printer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" transform="rotate(-90 10.5 8.5)"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.5 4.384V2.486a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2V14.5a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2v-2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M1.5 4.5h5.001v8H1.5a1 1 0 0 1-1-1v-6a1 1 0 0 1 1-1zm12 0h2a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1h-2"/><circle cx="9" cy="14" r="1" fill="currentColor"/></g>`),
		g.Group(children),
	)
}