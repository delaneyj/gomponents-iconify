package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignCenterHorizontally(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 6a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1h5v4.5a.5.5 0 0 0 1 0V9h5a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1H8V1.5a.5.5 0 0 0-1 0V6H2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}