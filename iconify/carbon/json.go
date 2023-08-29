package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Json(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M31 11v10h-2l-2-6v6h-2V11h2l2 6v-6h2zm-9.666 10h-2.667A1.668 1.668 0 0 1 17 19.334v-6.667A1.668 1.668 0 0 1 18.666 11h2.667A1.668 1.668 0 0 1 23 12.666v6.667A1.668 1.668 0 0 1 21.334 21zM19 19h2v-6h-2zm-5.666 2H9v-2h4v-2h-2a2.002 2.002 0 0 1-2-2v-2.334A1.668 1.668 0 0 1 10.666 11H15v2h-4v2h2a2.002 2.002 0 0 1 2 2v2.333A1.668 1.668 0 0 1 13.334 21zm-8 0H2.667A1.668 1.668 0 0 1 1 19.334V17h2v2h2v-8h2v8.334A1.668 1.668 0 0 1 5.333 21z"/>`),
		g.Group(children),
	)
}