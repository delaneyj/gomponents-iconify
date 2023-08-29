package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M751 105L448 408v104h128q27 0 45.5 19t18.5 45.5t-18.5 45T576 640H448v64h128q27 0 45.5 19t18.5 45.5t-18.5 45T576 832H448v128q0 27-18.5 45.5t-45 18.5t-45-18.5T321 960V832H193q-27 0-45.5-18.5t-18.5-45t18.5-45.5t45.5-19h128v-64H193q-27 0-45.5-18.5t-18.5-45t18.5-45.5t45.5-19h128V408L18 105Q0 87 0 61.5T18 18T61.5 0T105 18l279 279L664 18q18-18 43.5-18T751 18t18 43.5t-18 43.5z"/>`),
		g.Group(children),
	)
}