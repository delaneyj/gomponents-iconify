package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Todo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 3.5h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-10a2 2 0 0 1-2-2v-10a2 2 0 0 1 2-2z"/><g fill="currentColor" transform="translate(3 3)"><circle cx="7.5" cy="7.5" r="1" transform="matrix(-1 0 0 1 15 0)"/><circle cx="3.5" cy="7.5" r="1"/><circle cx="11.5" cy="7.5" r="1"/></g></g>`),
		g.Group(children),
	)
}