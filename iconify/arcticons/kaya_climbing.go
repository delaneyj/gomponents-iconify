package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KayaClimbing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".972" d="M6.279 6.151h37v37h-37z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m6.279 6.151l37 37"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".972" d="M6.28 43.15L18.1 17.972l25.18-11.82"/>`),
		g.Group(children),
	)
}