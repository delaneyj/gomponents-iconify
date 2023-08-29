package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanoramaVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m340 411l2 6q0 12-14 12H13q-13 0-13-12q0-3 1-6q35-95 35-195T1 21q-1-3-1-6Q0 3 13 3h315q13 0 13 12q0 3-1 6q-35 95-35 195q0 101 35 195zM54 387h233q-25-84-25-171t25-171H54q25 84 25 171T54 387z"/>`),
		g.Group(children),
	)
}