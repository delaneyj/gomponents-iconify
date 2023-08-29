package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linejoin="round" transform="translate(2 2)"><circle cx="8.5" cy="8.5" r="8" stroke-linecap="round"/><path stroke-linecap="round" d="m9.5 4.5l2 2v1a3 3 0 0 1-6 0v-1z"/><path d="M3.5 12V7.5a5 5 0 1 1 10 0V12"/><path stroke-linecap="round" d="M14.5 13.404c-.662-2.273-3.2-2.93-6-2.93c-2.727 0-5.27.774-6 2.93"/></g>`),
		g.Group(children),
	)
}