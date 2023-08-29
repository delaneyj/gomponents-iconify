package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Door(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" transform="translate(4 3)"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5.5h8a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-8a2 2 0 0 1-2-2v-10a2 2 0 0 1 2-2z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m10.202 14.5l-3.645-1.948A2 2 0 0 1 5.5 10.788V4.212a2 2 0 0 1 1.057-1.764L10.202.5"/><circle cx="7.5" cy="7.5" r="1" fill="currentColor"/></g>`),
		g.Group(children),
	)
}