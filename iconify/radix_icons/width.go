package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Width(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.818 4.682a.45.45 0 0 1 0 .636L3.086 7.05h8.828l-1.732-1.732a.45.45 0 0 1 .636-.636l2.5 2.5a.45.45 0 0 1 0 .636l-2.5 2.5a.45.45 0 0 1-.636-.636l1.732-1.732H3.086l1.732 1.732a.45.45 0 1 1-.636.636l-2.5-2.5a.45.45 0 0 1 0-.636l2.5-2.5a.45.45 0 0 1 .636 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}