package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Replika(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m21.438 14.094l4.402 1.756l3.734-2.32l4.2 2.372l3.056-3.234"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4.5c8.784 0 17.319 11.402 17.319 21.75A17.264 17.264 0 0 1 24 43.5m0-39c-8.784 0-17.319 11.402-17.319 21.75A17.264 17.264 0 0 0 24 43.5"/>`),
		g.Group(children),
	)
}