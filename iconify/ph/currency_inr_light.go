package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyInrLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M206 80a6 6 0 0 1-6 6h-34.31a59.36 59.36 0 0 1 .31 6a58.07 58.07 0 0 1-58 58H87.52L164 219.56a6 6 0 0 1-8.08 8.88l-88-80A6 6 0 0 1 72 138h36a46.06 46.06 0 0 0 46-46a47.61 47.61 0 0 0-.4-6H72a6 6 0 0 1 0-12h78.33A46.08 46.08 0 0 0 108 46H72a6 6 0 0 1 0-12h128a6 6 0 0 1 0 12h-56.73a58.25 58.25 0 0 1 19.86 28H200a6 6 0 0 1 6 6Z"/>`),
		g.Group(children),
	)
}