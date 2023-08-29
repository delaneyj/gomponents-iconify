package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsClockwiseLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M196.24 187.76a6 6 0 0 1 0 8.48C195.19 197.29 170 222 128 222c-39.66 0-67.59-25.75-82-43.26V208a6 6 0 0 1-12 0v-48a6 6 0 0 1 6-6h48a6 6 0 0 1 0 12H51.35c11.41 15.11 38.23 44 76.65 44c37.09 0 59.54-22 59.76-22.24a6 6 0 0 1 8.48 0ZM216 42a6 6 0 0 0-6 6v29.26C195.59 59.75 167.66 34 128 34c-42 0-67.19 24.71-68.24 25.76a6 6 0 0 0 8.48 8.48C68.46 68 90.91 46 128 46c38.42 0 65.24 28.89 76.65 44H168a6 6 0 0 0 0 12h48a6 6 0 0 0 6-6V48a6 6 0 0 0-6-6Z"/>`),
		g.Group(children),
	)
}