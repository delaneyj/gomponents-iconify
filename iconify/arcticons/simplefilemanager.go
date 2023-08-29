package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Simplefilemanager(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.66 12.5v-2.21h0a1.93 1.93 0 0 1 1.94-1.94h8.85c2 0 5.56 3.69 7.32 3.78h15.1a1.63 1.63 0 0 1 1.63 1.63v19.79h0a1.94 1.94 0 0 1-2 1.94h-2.16V17.92a1.62 1.62 0 0 0-1.62-1.63H22.61c-1.75-.1-5.3-3.79-7.32-3.79H6.45a1.94 1.94 0 0 0-1.95 1.94h0v23.27a1.94 1.94 0 0 0 1.94 1.94h31a1.93 1.93 0 0 0 1.94-1.94h0V21"/>`),
		g.Group(children),
	)
}