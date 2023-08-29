package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yalp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.76 8.74L21.77 24l-5-5l17.9-9.95a6.11 6.11 0 0 1 1.73-.64a2.06 2.06 0 0 1 1.36.33Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.76 8.74a3.59 3.59 0 0 1 .89 2.85v24.82a3.54 3.54 0 0 1-.9 2.83L21.77 24l16-15.26Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m16.81 19l5 5l-5 5l-5.88-3.26c-2.11-1.17-2.1-2.4 0-3.56L16.81 19Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m21.77 24l16 15.24c-.78.71-2.1.25-3.07-.29L16.81 29Z"/>`),
		g.Group(children),
	)
}