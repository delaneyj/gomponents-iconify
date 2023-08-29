package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GifLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M142 72v112a6 6 0 0 1-12 0V72a6 6 0 0 1 12 0Zm82-6h-48a6 6 0 0 0-6 6v112a6 6 0 0 0 12 0v-50h34a6 6 0 0 0 0-12h-34V78h42a6 6 0 0 0 0-12ZM96 122H72a6 6 0 0 0 0 12h18v18a26 26 0 0 1-52 0v-48a26 26 0 0 1 26-26c12.07 0 23.33 8.38 26.19 19.5a6 6 0 1 0 11.62-3C97.56 78 81.66 66 64 66a38 38 0 0 0-38 38v48a38 38 0 0 0 76 0v-24a6 6 0 0 0-6-6Z"/>`),
		g.Group(children),
	)
}