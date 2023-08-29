package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SewingPin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6 3.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Zm2 2.45a2.5 2.5 0 1 0-1 0v7.55a.5.5 0 0 0 1 0V5.95Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}