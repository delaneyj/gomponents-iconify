package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Openlauncher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="8" x="5.5" y="34.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="4"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.5 5.5h29a4 4 0 0 1 4 4h0a4 4 0 0 1-4 4h-25"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 38.5v-29a4 4 0 0 1 4-4h0a4 4 0 0 1 4 4v25"/>`),
		g.Group(children),
	)
}