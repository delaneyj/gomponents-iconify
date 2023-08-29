package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableOffsetTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 3H6.25A3.25 3.25 0 0 0 3 6.25V8.5h11V3Zm1.5 5.5H21V6.25A3.25 3.25 0 0 0 17.75 3H15.5v5.5ZM21 10H10v4h11v-4ZM8.5 10H3v4h5.5v-4ZM3 17.75V15.5h11V21H6.25A3.25 3.25 0 0 1 3 17.75ZM15.5 21v-5.5H21v2.25A3.25 3.25 0 0 1 17.75 21H15.5Z"/>`),
		g.Group(children),
	)
}