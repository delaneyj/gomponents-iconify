package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func P(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M0 0v1408h384V896h320q198 0 322.5-119T1151 448t-124.5-329T704 0H0zm384 640V256h192q89 0 140.5 48.5T768 448t-53.5 143.5T576 640H384z"/>`),
		g.Group(children),
	)
}