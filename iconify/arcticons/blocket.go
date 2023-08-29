package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blocket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.5 4.5h-27a2 2 0 0 0-2 2v35a2 2 0 0 0 2 2h27a2 2 0 0 0 2-2v-35a2 2 0 0 0-2-2ZM12 7V3m6 4V3m6 4V3m6 4V3m6 4V3"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18 26.1a6 6 0 0 1 12 0V30a6 6 0 0 1-12 0m0 6V12"/>`),
		g.Group(children),
	)
}