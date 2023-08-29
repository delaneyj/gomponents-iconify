package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IphoneLandscape(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" transform="translate(3 5)"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5.5h10a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2h-10a2 2 0 0 1-2-2v-6a2 2 0 0 1 2-2z"/><circle cx="11.5" cy="5.5" r="1" fill="currentColor"/></g>`),
		g.Group(children),
	)
}