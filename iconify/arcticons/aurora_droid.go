package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AuroraDroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.55 14.64a1.9 1.9 0 1 1-1.91 1.9a1.9 1.9 0 0 1 1.91-1.9Zm-11.08 0a1.9 1.9 0 1 1-1.9 1.9a1.91 1.91 0 0 1 1.9-1.9Zm17.22 7.54H12.31a1.87 1.87 0 0 0-1.87 1.87v5.89a13.56 13.56 0 0 0 27.12 0v-5.89a1.87 1.87 0 0 0-1.87-1.87Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.01 34.55h3.98L24 31.04l-1.99 3.51z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.35 37.09a1.7 1.7 0 0 1-1.46.84h-9.78a1.69 1.69 0 0 1-1.47-2.52l4.89-8.62a1.69 1.69 0 0 1 2.94 0l4.89 8.62a1.71 1.71 0 0 1-.01 1.68ZM24 6.75h0a13.56 13.56 0 0 1 13.56 13.56v0a1.87 1.87 0 0 1-1.87 1.87H12.31a1.87 1.87 0 0 1-1.87-1.87v0A13.56 13.56 0 0 1 24 6.75ZM11.76 4.87l4.4 4.39m20.07-4.39l-4.39 4.4"/>`),
		g.Group(children),
	)
}