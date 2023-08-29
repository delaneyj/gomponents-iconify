package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Entel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.629 30.077a7.13 7.13 0 0 1-6.262 3.684A7.389 7.389 0 0 1 13 26.394v-4.788a7.367 7.367 0 0 1 14.734 0v2.578H13M31.513 16.6a9.596 9.596 0 0 1 0 14.8"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2 2 0 0 1-2-2v-33a2 2 0 0 1 2-2h33a2 2 0 0 1 2 2v33a2 2 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}