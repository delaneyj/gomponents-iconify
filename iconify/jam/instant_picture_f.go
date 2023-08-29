package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InstantPictureF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16 11.252l-1.454-2.24a3 3 0 0 0-5-.304L5.575 14H0V2a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v9.252zM8.076 14l3.07-4.092a1 1 0 0 1 1.666.101L15.284 14H8.076zM16 16v2a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2v-2h16zM6 9a3 3 0 1 0 0-6a3 3 0 0 0 0 6z"/>`),
		g.Group(children),
	)
}