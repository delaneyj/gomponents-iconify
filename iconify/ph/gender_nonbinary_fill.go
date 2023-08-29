package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GenderNonbinaryFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 24H56a16 16 0 0 0-16 16v176a16 16 0 0 0 16 16h144a16 16 0 0 0 16-16V40a16 16 0 0 0-16-16Zm-72 184a52 52 0 0 1-8-103.38v-20.8L91 95.43a8 8 0 0 1-6-14.86L106.46 72L85 63.43a8 8 0 0 1 6-14.86l37 14.81l37-14.81a8 8 0 1 1 6 14.86L149.54 72L171 80.57a8 8 0 0 1-6 14.86l-29-11.61v20.8A52 52 0 0 1 128 208Zm36-52a36 36 0 1 1-36-36a36 36 0 0 1 36 36Z"/>`),
		g.Group(children),
	)
}