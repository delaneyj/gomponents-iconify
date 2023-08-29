package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dmss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m4.474 26.318l17.112 17.196a3.346 3.346 0 0 0 4.732.012l17.196-17.112a3.346 3.346 0 0 0 .012-4.732L26.414 4.486a3.346 3.346 0 0 0-4.732-.012L4.486 21.586a3.346 3.346 0 0 0-.012 4.732Z"/><circle cx="24" cy="24" r="11.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.935 23.13a3.065 3.065 0 0 1-3.066-3.065c0-1.128.617-2.104 1.525-2.636A6.969 6.969 0 0 0 24 17a7 7 0 1 0 7 7a6.94 6.94 0 0 0-.43-2.394a3.053 3.053 0 0 1-2.635 1.525Z"/>`),
		g.Group(children),
	)
}