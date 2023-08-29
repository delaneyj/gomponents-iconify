package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderDotted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 6.625a.875.875 0 1 0 0 1.75a.875.875 0 0 0 0-1.75Zm4 0a.875.875 0 1 0 0 1.75a.875.875 0 0 0 0-1.75Zm4 0a.875.875 0 1 0 0 1.75a.875.875 0 0 0 0-1.75Zm3.125.875a.875.875 0 1 1 1.75 0a.875.875 0 0 1-1.75 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}