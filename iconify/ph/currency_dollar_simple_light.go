package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyDollarSimpleLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M198 168a46.06 46.06 0 0 1-46 46h-18v18a6 6 0 0 1-12 0v-18h-18a46.06 46.06 0 0 1-46-46a6 6 0 0 1 12 0a34 34 0 0 0 34 34h48a34 34 0 0 0 0-68h-40a46 46 0 0 1 0-92h10V24a6 6 0 0 1 12 0v18h10a46.06 46.06 0 0 1 46 46a6 6 0 0 1-12 0a34 34 0 0 0-34-34h-32a34 34 0 0 0 0 68h40a46.06 46.06 0 0 1 46 46Z"/>`),
		g.Group(children),
	)
}