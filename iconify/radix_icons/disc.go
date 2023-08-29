package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Disc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.5.877a6.623 6.623 0 1 0 0 13.246A6.623 6.623 0 0 0 7.5.877ZM1.827 7.5a5.673 5.673 0 1 1 11.346 0a5.673 5.673 0 0 1-11.346 0Zm6.546 0a.873.873 0 1 1-1.745 0a.873.873 0 0 1 1.745 0Zm.95 0a1.823 1.823 0 1 1-3.645 0a1.823 1.823 0 0 1 3.645 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}