package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyRupee(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24 7V5H8v2h5.5a4.49 4.49 0 0 1 4.45 4H8v2h9.95a4.49 4.49 0 0 1-4.45 4H8v2.345L17.617 28l1.338-1.486L10.606 19H13.5a6.504 6.504 0 0 0 6.475-6H24v-2h-4.025a6.472 6.472 0 0 0-1.795-4Z"/>`),
		g.Group(children),
	)
}