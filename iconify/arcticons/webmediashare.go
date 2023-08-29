package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Webmediashare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.5 39.5h14a2 2 0 0 0 2-2v-27a2 2 0 0 0-2-2h-35a2 2 0 0 0-2 2v6m0 16h0a7 7 0 0 1 7 7v0h0h-7h0v-7h0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.5 39.5h0a19 19 0 0 0-19-19h0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.3 39.5h0A12.8 12.8 0 0 0 4.5 26.7h0"/>`),
		g.Group(children),
	)
}