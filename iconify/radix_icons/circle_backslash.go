package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleBackslash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.5.877a6.623 6.623 0 1 0 0 13.246A6.623 6.623 0 0 0 7.5.877ZM3.858 3.151a5.673 5.673 0 0 1 7.992 7.992L3.857 3.15Zm-.707.707a5.673 5.673 0 0 0 7.992 7.992L3.15 3.857Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}