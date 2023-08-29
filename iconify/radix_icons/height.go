package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Height(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.181 1.682a.45.45 0 0 1 .637 0l2.5 2.5a.45.45 0 0 1-.637.636L7.95 3.086v8.828l1.731-1.732a.45.45 0 0 1 .637.636l-2.5 2.5a.45.45 0 0 1-.637 0l-2.5-2.5a.45.45 0 0 1 .637-.636l1.732 1.732V3.086L5.317 4.818a.45.45 0 0 1-.637-.636l2.5-2.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}