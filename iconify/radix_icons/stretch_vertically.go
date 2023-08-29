package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StretchVertically(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 .5a.5.5 0 0 1 .5-.5h12a.5.5 0 0 1 0 1h-12A.5.5 0 0 1 1 .5ZM9 14V1H6v13H1.5a.5.5 0 0 0 0 1h12a.5.5 0 0 0 0-1H9Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}