package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Libretranslater(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.4 6.45v35.1a1.95 1.95 0 0 0 1.95 1.95h2.376v-39H10.35A1.95 1.95 0 0 0 8.4 6.45Zm4.326-1.95v39H37.65a1.95 1.95 0 0 0 1.95-1.95V6.45a1.95 1.95 0 0 0-1.95-1.95Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.332 24.399v6.862h6.862m2.014 4.214l2.813-8.452m2.697 8.477l-2.697-8.477m1.795 5.641h-3.673M16.608 14.372h9.447M21.332 12.5v1.872m2.644 0c0 2.584-3 6.863-6.001 7.516"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.45 16.719c.297 1.782 3.387 4.753 6.135 5.17m-5.021 4.277l1.768-1.767l1.767 1.767m3.328 3.328l1.767 1.767l-1.767 1.768"/>`),
		g.Group(children),
	)
}