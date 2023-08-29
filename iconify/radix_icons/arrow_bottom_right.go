package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBottomRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.646 3.646a.5.5 0 0 0 0 .708L10.293 11H6a.5.5 0 0 0 0 1h5.5a.5.5 0 0 0 .5-.5V6a.5.5 0 0 0-1 0v4.293L4.354 3.646a.5.5 0 0 0-.708 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}