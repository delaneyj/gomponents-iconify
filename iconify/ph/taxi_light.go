package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TaxiLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M240 114h-12.52l-28-48.95a14 14 0 0 0-12.15-7h-23.3L151.58 26.8a13.93 13.93 0 0 0-13-8.8h-21.16a13.93 13.93 0 0 0-13 8.8L91.94 58h-23.3a14 14 0 0 0-12.15 7l-28 49H16a6 6 0 0 0 0 12h10v82a14 14 0 0 0 14 14h24a14 14 0 0 0 14-14v-18h100v18a14 14 0 0 0 14 14h24a14 14 0 0 0 14-14v-82h10a6 6 0 0 0 0-12ZM115.56 31.26a2 2 0 0 1 1.86-1.26h21.16a2 2 0 0 1 1.86 1.26L151.14 58h-46.28ZM66.91 71a2 2 0 0 1 1.73-1h118.72a2 2 0 0 1 1.73 1l24.57 43H42.34ZM66 208a2 2 0 0 1-2 2H40a2 2 0 0 1-2-2v-18h28Zm150 2h-24a2 2 0 0 1-2-2v-18h28v18a2 2 0 0 1-2 2Zm2-32H38v-52h180ZM58 152a6 6 0 0 1 6-6h16a6 6 0 0 1 0 12H64a6 6 0 0 1-6-6Zm112 0a6 6 0 0 1 6-6h16a6 6 0 0 1 0 12h-16a6 6 0 0 1-6-6Z"/>`),
		g.Group(children),
	)
}