package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Italic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M13.5 1.25V0H4.71v1.25h3.67l-2.02 13.5H2.5V16h8.79v-1.25H7.62l2.02-13.5h3.86z"/>`),
		g.Group(children),
	)
}