package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyKztFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 24a104 104 0 1 0 104 104A104.11 104.11 0 0 0 128 24Zm48 96h-40v72a8 8 0 0 1-16 0v-72H80a8 8 0 0 1 0-16h96a8 8 0 0 1 0 16Zm0-32H80a8 8 0 0 1 0-16h96a8 8 0 0 1 0 16Z"/>`),
		g.Group(children),
	)
}