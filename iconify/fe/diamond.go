package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Diamond(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feDiamond0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feDiamond1" fill="currentColor" fill-rule="nonzero"><path id="feDiamond2" d="m12 17.876l6.626-7.952L16.164 5H7.836L5.374 9.924L12 17.876ZM6.6 3h10.8l3.6 7.2L12 21L3 10.2L6.6 3Z"/></g></g>`),
		g.Group(children),
	)
}