package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyJpyFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 24a104 104 0 1 0 104 104A104.11 104.11 0 0 0 128 24Zm62 53.27L145.63 128H160a8 8 0 0 1 0 16h-24v16h24a8 8 0 0 1 0 16h-24v16a8 8 0 0 1-16 0v-16H96a8 8 0 0 1 0-16h24v-16H96a8 8 0 0 1 0-16h14.37L66 77.27a8 8 0 0 1 12-10.54l50 57.12l50-57.12a8 8 0 1 1 12 10.54Z"/>`),
		g.Group(children),
	)
}