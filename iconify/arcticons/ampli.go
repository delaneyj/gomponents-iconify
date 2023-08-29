package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ampli(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M24 2.5A21.5 21.5 0 1 1 2.5 24A21.503 21.503 0 0 1 24 2.5Z"/><rect width="17.949" height="21" x="14.025" y="13.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="8.975"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.974 34.5v-21"/>`),
		g.Group(children),
	)
}