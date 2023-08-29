package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Repainter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.538 24.326L38.54 5.5h-3.726v10.213h-4.412L30.33 5.5l.04 5.814l-2.895 2.774h-1.888v-1.907H23.56V5.5H9.461v18.826m-.001 0v4.154a3.115 3.115 0 0 0 3.116 3.116h7.789V42.5h7.27V31.596h7.788a3.115 3.115 0 0 0 3.116-3.116v-4.154H9.46Z"/>`),
		g.Group(children),
	)
}