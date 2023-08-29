package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RockOn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="m19 16l.87-11.735a2.102 2.102 0 0 0-4.181-.433L15 9m-7 2l-.713-4.279a2.06 2.06 0 0 0-4.083.525L4 16m8-3v-1.5a2 2 0 1 0-4 0V15"/><path stroke-linecap="round" stroke-linejoin="round" d="M16 13v-2a2 2 0 1 0-4 0v2"/><path d="M19 16c-.536 4-3.358 6-7.5 6C7.358 22 4 20 4 16"/><path stroke-linecap="round" stroke-linejoin="round" d="M13.692 17H11a2 2 0 1 1 0-4h4c2.21 0 4.5 2 3.5 5"/></g>`),
		g.Group(children),
	)
}