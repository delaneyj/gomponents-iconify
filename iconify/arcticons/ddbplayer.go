package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ddbplayer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.35 24h11.3a6.75 6.75 0 0 0 6.74-6.75h0a6.75 6.75 0 0 0-6.74-6.75h-14l2.74 2.85M16.35 24h11.3a6.75 6.75 0 0 1 6.74 6.75h0a6.75 6.75 0 0 1-6.74 6.75h-11.3"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.35 10.5h-2.74l2.74 2.85m0 10.65h-2.74l2.74 2.85m0-2.85V13.35"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.35 37.5h-2.74l2.74-2.85v-7.8"/>`),
		g.Group(children),
	)
}