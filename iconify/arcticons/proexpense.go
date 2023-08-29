package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Proexpense(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="square" d="M36.3 11.1v25m-.1.3h-25m-.1 0v-25m.2-.3h25"/><path fill="none" stroke="currentColor" stroke-linecap="round" d="M6.6 11.3A3.4 3.4 0 0 1 8 5.4a3.28 3.28 0 0 1 3.3 1.2M6.6 36.4v-25m4.7 29.7a3.4 3.4 0 0 1-5.9-1.4a3.28 3.28 0 0 1 1.2-3.3m29.7 4.7h-25m29.8-4.6a3.32 3.32 0 0 1 1.1 3.4a3.51 3.51 0 0 1-2.5 2.5a3.28 3.28 0 0 1-3.3-1.2m4.7-29.7v25M36.4 6.7a3.32 3.32 0 0 1 3.4-1.1a3.51 3.51 0 0 1 2.5 2.5a3.28 3.28 0 0 1-1.2 3.3M11.4 6.7h25"/>`),
		g.Group(children),
	)
}