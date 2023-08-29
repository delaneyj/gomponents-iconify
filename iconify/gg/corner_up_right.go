package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerUpRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13.372 14.722l1.412 1.417l6.377-6.35l-6.35-6.378l-1.417 1.411l3.776 3.793l-10.305-.022a4 4 0 0 0-4.009 3.991l-.017 8l2 .005l.017-8a2 2 0 0 1 2.004-1.996l10.636.023l-4.124 4.106Z"/>`),
		g.Group(children),
	)
}