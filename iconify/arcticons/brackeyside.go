package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Brackeyside(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m16.11 20.42l-4.52-4.51L3.5 24L24 44.5h0l8.09-8.09l-4.51-4.52L24 35.47L12.53 24l3.58-3.58zm15.78 7.16l4.52 4.51L44.5 24L24 3.5h0l-8.09 8.09l4.51 4.52L24 12.53L35.47 24l-3.58 3.58z"/>`),
		g.Group(children),
	)
}