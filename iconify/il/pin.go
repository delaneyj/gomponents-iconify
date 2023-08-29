package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M531 99q45 46 68 103t23 117t-23 117t-68 103L327 743q-7 7-16 7t-16-7L91 539q-45-45-68-103T0 319t23-117T91.5 99T194 31T311 8t117 23t103 68zM383 391q15-15 22-34t7-38t-7-38t-22-34t-34-22t-38-7t-38 7t-34 22t-22 34t-8 38t8 38t22 34t34 22t38 8t38-8t34-22z"/>`),
		g.Group(children),
	)
}