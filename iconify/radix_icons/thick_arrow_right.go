package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThickArrowRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14 7.5a.5.5 0 0 1-.21.407l-7 5A.5.5 0 0 1 6 12.5V10H3.5a.5.5 0 0 1-.5-.5v-4a.5.5 0 0 1 .5-.5H6V2.5a.5.5 0 0 1 .79-.407l7 5A.5.5 0 0 1 14 7.5ZM7 3.472V5.5a.5.5 0 0 1-.5.5H4v3h2.5a.5.5 0 0 1 .5.5v2.028L12.64 7.5L7 3.472Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}