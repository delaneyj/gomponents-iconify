package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignCenterVertically(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7 1a1 1 0 0 0-1 1v5H1.5a.5.5 0 0 0 0 1H6v5a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1V8h4.5a.5.5 0 0 0 0-1H9V2a1 1 0 0 0-1-1H7Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}