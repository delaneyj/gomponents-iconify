package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleNewsstand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.39 8.4v8.94a2 2 0 0 1-2 1.95H4.5v18.36a2 2 0 0 0 2 2h35.1a2 2 0 0 0 2-2v-27.3a2 2 0 0 0-2-1.95ZM13 24.7a4.36 4.36 0 0 1 0 8.72a4.28 4.28 0 0 1-4.21-4.36h0A4.28 4.28 0 0 1 13 24.7Zm2.39-16.3L4.5 19.29m16.21 14.13h17.61M20.71 14.58h17.61M20.71 24h17.61"/>`),
		g.Group(children),
	)
}