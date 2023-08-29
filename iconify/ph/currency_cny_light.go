package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyCnyLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M58 64a6 6 0 0 1 6-6h128a6 6 0 0 1 0 12H64a6 6 0 0 1-6-6Zm158 106a6 6 0 0 0-6 6v18h-34a18 18 0 0 1-18-18v-50h50a6 6 0 0 0 0-12H48a6 6 0 0 0 0 12h50v10a58.07 58.07 0 0 1-58 58a6 6 0 0 0 0 12a70.08 70.08 0 0 0 70-70v-10h36v50a30 30 0 0 0 30 30h40a6 6 0 0 0 6-6v-24a6 6 0 0 0-6-6Z"/>`),
		g.Group(children),
	)
}