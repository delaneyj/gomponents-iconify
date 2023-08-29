package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Instagrabber(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.05" d="M5.5 40.55a2 2 0 0 0 2 2h33.1a2 2 0 0 0 2-2V7.45a2 2 0 0 0-2-1.95H7.45a2 2 0 0 0-1.95 2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.05" d="M24 35.44a1.2 1.2 0 0 1-1-.44l-7.55-9.84a1.23 1.23 0 0 1-.13-1.28a1.22 1.22 0 0 1 1.1-.68h3l-6.63-8.64a1.24 1.24 0 0 1-.13-1.28a1.22 1.22 0 0 1 1.1-.68h20.46a1.22 1.22 0 0 1 1.1.68a1.24 1.24 0 0 1-.13 1.28l-6.63 8.64h3a1.22 1.22 0 0 1 1.1.68a1.23 1.23 0 0 1-.13 1.28L25 35a1.2 1.2 0 0 1-1 .44Z"/>`),
		g.Group(children),
	)
}