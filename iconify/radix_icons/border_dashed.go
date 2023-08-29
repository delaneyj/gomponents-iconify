package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderDashed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5A.5.5 0 0 1 .5 7H3a.5.5 0 0 1 0 1H.5a.5.5 0 0 1-.5-.5Zm5.75 0a.5.5 0 0 1 .5-.5h2.5a.5.5 0 0 1 0 1h-2.5a.5.5 0 0 1-.5-.5ZM12 7a.5.5 0 0 0 0 1h2.5a.5.5 0 0 0 0-1H12Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}