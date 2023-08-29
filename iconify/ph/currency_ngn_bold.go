package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyNgnBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 116h-12V46a12 12 0 0 0-24 0v70h-46.14l-60.4-77.38A12 12 0 0 0 52 46v70H40a12 12 0 0 0 0 24h12v70a12 12 0 0 0 24 0v-70h46.14l60.4 77.38A12 12 0 0 0 204 210v-70h12a12 12 0 0 0 0-24Zm-140 0V80.88L103.41 116Zm104 59.12L152.59 140H180Z"/>`),
		g.Group(children),
	)
}