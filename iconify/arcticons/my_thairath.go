package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MyThairath(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m36.285 19.483l-.88-1.313c-.744-1.11-1.701-1.767-2.998-1.767h-8.445c-1.698 0-3.563 1.42-4.09 3.035l-7.06 21.604m-7.312 0h13.752M5.965 6.958l.928 1.355C7.767 9.588 7.935 9.7 9.377 9.7h27.48c1.92 0 2.16.098 3.183 1.567l2.46 3.535"/>`),
		g.Group(children),
	)
}