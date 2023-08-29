package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blade(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M8 14v3H6a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h2v3a2 2 0 0 0 2 2h27a2 2 0 0 0 2-2v-3h3a2 2 0 0 0 2-2V19a2 2 0 0 0-2-2h-3v-3a2 2 0 0 0-2-2H10a2 2 0 0 0-2 2Zm4 10h24"/><circle cx="24" cy="24" r="4" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M32 29V19M16 29V19"/></g>`),
		g.Group(children),
	)
}