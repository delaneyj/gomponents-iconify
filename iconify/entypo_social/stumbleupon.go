package entypo_social

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stumbleupon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m11.051 8.059l1.365.66l2.059-.66V6.865C14.475 4.402 12.467 2.4 10 2.4S5.525 4.402 5.525 6.865v6.27a1.052 1.052 0 0 1-2.102 0V10.51H0v2.625a4.475 4.475 0 0 0 8.949 0v-6.27a1.052 1.052 0 0 1 2.103 0v1.194zm5.525 2.451v2.625a1.051 1.051 0 0 1-2.102 0v-2.678l-2.059.658l-1.365-.658v2.678a4.476 4.476 0 0 0 8.95 0V10.51h-3.424z"/>`),
		g.Group(children),
	)
}