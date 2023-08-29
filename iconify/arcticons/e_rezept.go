package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ERezept(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="44.974" height="17.75" x="1.513" y="15.125" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="8.875" transform="rotate(-45 24 24)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m36.551 24l3.35 3.35a8.875 8.875 0 0 1 0 12.55h0a8.875 8.875 0 0 1-12.552 0L8.1 20.65a8.875 8.875 0 0 1 0-12.55h0a8.875 8.875 0 0 1 12.552 0L24 11.448"/>`),
		g.Group(children),
	)
}