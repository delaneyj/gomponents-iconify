package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.628 14.722l-1.412 1.417L2.84 9.79l6.35-6.378l1.417 1.411L6.83 8.615l10.305-.022a4 4 0 0 1 4.009 3.991l.017 8l-2 .005l-.017-8a2 2 0 0 0-2.004-1.996l-10.636.023l4.124 4.106Z"/>`),
		g.Group(children),
	)
}