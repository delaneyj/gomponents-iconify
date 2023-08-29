package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SteamChat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="31.058" cy="18.366" r="7.957" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="17.392" cy="32.081" r="5.62" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m11.852 33.025l-8.668-3.874M2.56 22.39l11.523 5.15"/><path fill="none" stroke="currentColor" stroke-linecap="round" d="m23.1 18.366l-5.708 8.095m5.62 5.62l8.046-5.757"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5A21.5 21.5 0 0 0 2.5 24a23.139 23.139 0 0 0 2.665 10.297L2.5 45.5l11.174-2.658A21.494 21.494 0 1 0 24 2.5Z"/>`),
		g.Group(children),
	)
}