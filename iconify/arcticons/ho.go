package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ho(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><rect width="6.85" height="9.076" x="19.65" y="29.424" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.425"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.5 24.801V38.5m0-5.65a3.425 3.425 0 0 1 3.425-3.426h0a3.425 3.425 0 0 1 3.425 3.425V38.5"/><circle cx="29.009" cy="38.255" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}