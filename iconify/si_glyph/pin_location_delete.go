package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinLocationDelete(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.475.031c-3.007 0-5.443 2.512-5.443 5.609c0 1.584 5.443 10.275 5.443 10.275s5.441-8.609 5.441-10.275c0-3.097-2.437-5.609-5.441-5.609zm3.474 7.867L9.914 8.932L7.5 6.518L5.086 8.932L4.051 7.898l2.414-2.414L4.051 3.07l1.035-1.035L7.5 4.449l2.414-2.414l1.035 1.035l-2.415 2.414l2.415 2.414z"/>`),
		g.Group(children),
	)
}