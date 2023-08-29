package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PokemonHome(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 35.671V35.5C5.5 25.282 13.783 17 24 17s18.5 8.282 18.5 18.5v.171"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.561 37.68a24.746 24.746 0 0 1 12.854-7.357M42.5 37.68a24.746 24.746 0 0 0-12.854-7.357"/><circle cx="24" cy="29.366" r="5.237" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><circle cx="24" cy="29.366" r="2.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}