package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Francetv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2ZM18.555 17.899h8.332"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m36 17.899l-4.557 12.075l-4.556-12.075m-6.054-3.873v13.67a2.152 2.152 0 0 0 2.279 2.278h.683"/><circle cx="14.636" cy="23.937" r="2.636" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}