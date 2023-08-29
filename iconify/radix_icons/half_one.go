package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HalfOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.5.877a6.623 6.623 0 1 0 0 13.246A6.623 6.623 0 0 0 7.5.877Zm-.5.972A5.673 5.673 0 0 0 7 13.15V1.849ZM8 13.15A5.673 5.673 0 0 0 8 1.849V13.15Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}