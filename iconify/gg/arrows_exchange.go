package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsExchange(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.993 12.984a1 1 0 0 0-.531 1.848L7.15 17.52a1 1 0 1 0 1.414-1.415l-1.121-1.12h7.55a1 1 0 0 0 0-2h-10Zm14.014-1.968a1 1 0 0 0 .531-1.848L16.85 6.48a1 1 0 0 0-1.414 1.415l1.121 1.12h-7.55a1 1 0 0 0 0 2h10Z"/>`),
		g.Group(children),
	)
}