package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageRotate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.25 1.02c5.1 0 8.75 4.04 8.75 9s-3.65 9-8.75 9a9.32 9.32 0 0 1-7.68-3.99l2.59-1.52c1.1 1.5 2.86 2.51 4.84 2.51c3.3 0 6-2.79 6-6s-2.7-6-6-6c-1.97 0-3.72 1-4.82 2.49L7 8.02l-6 2v-7L2.89 4.6a9.315 9.315 0 0 1 7.36-3.58z"/>`),
		g.Group(children),
	)
}