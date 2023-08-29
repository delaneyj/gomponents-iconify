package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stocard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.68 19.85a1.65 1.65 0 1 0-2.34 0a1.66 1.66 0 0 0 2.34 0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.46 13.83L17.65 9a2.09 2.09 0 0 0-2.95 0l-9.59 9.61a2.09 2.09 0 0 0 0 2.95l6.36 6.36L30.36 9a2.09 2.09 0 0 1 3 0l9.58 9.58a2.09 2.09 0 0 1 0 2.95L25.48 39a2.11 2.11 0 0 1-3 0l-8.79-8.79"/>`),
		g.Group(children),
	)
}