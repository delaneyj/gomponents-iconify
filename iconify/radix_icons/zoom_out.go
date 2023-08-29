package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomOut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.5 10a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7Zm0 1a4.48 4.48 0 0 0 2.809-.984l2.837 2.838a.5.5 0 0 0 .708-.708L10.016 9.31A4.5 4.5 0 1 0 6.5 11ZM4.75 6a.5.5 0 0 0 0 1h3.5a.5.5 0 0 0 0-1h-3.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}