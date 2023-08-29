package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloseXs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.28 3.22a.75.75 0 0 0-1.06 1.06L4.94 6L3.22 7.72a.75.75 0 0 0 1.06 1.06L6 7.06l1.72 1.72a.75.75 0 0 0 1.06-1.06L7.06 6l1.72-1.72a.75.75 0 0 0-1.06-1.06L6 4.94L4.28 3.22Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}