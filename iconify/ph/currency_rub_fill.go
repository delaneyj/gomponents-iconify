package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyRubFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M176 108a28 28 0 0 1-28 28h-36V80h36a28 28 0 0 1 28 28Zm56 20A104 104 0 1 1 128 24a104.11 104.11 0 0 1 104 104Zm-40-20a44.05 44.05 0 0 0-44-44h-44a8 8 0 0 0-8 8v64H80a8 8 0 0 0 0 16h16v16H80a8 8 0 0 0 0 16h16v16a8 8 0 0 0 16 0v-16h40a8 8 0 0 0 0-16h-40v-16h36a44.05 44.05 0 0 0 44-44Z"/>`),
		g.Group(children),
	)
}