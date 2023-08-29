package fa_6_brands

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vuejs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M356.9 64.3H280l-56 88.6l-48-88.6H0L224 448L448 64.3h-91.1zm-301.2 32h53.8L224 294.5L338.4 96.3h53.8L224 384.5L55.7 96.3z"/>`),
		g.Group(children),
	)
}