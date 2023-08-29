package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SewingPinFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10 3.5a2.5 2.5 0 0 1-2 2.45v7.55a.5.5 0 0 1-1 0V5.95a2.5 2.5 0 1 1 3-2.45Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}