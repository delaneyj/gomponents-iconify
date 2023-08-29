package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill-rule="evenodd" d="M7.5 8l-5 5L1 11.5L4.75 8L1 4.5L2.5 3l5 5z" fill="currentColor"/>`),
		g.Group(children),
	)
}