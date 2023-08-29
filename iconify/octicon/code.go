package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Code(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill-rule="evenodd" d="M9.5 3L8 4.5L11.5 8L8 11.5L9.5 13L14 8L9.5 3zm-5 0L0 8l4.5 5L6 11.5L2.5 8L6 4.5L4.5 3z" fill="currentColor"/>`),
		g.Group(children),
	)
}