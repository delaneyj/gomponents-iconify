package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Face(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.5.877a6.623 6.623 0 1 0 0 13.246A6.623 6.623 0 0 0 7.5.877ZM1.827 7.5a5.673 5.673 0 1 1 11.346 0a5.673 5.673 0 0 1-11.346 0Zm3.21 1.714a.5.5 0 1 0-.82.572A3.996 3.996 0 0 0 7.5 11.5c1.36 0 2.56-.679 3.283-1.714a.5.5 0 0 0-.82-.572A2.996 2.996 0 0 1 7.5 10.5a2.996 2.996 0 0 1-2.463-1.286Zm.338-2.364a.875.875 0 1 0 0-1.75a.875.875 0 0 0 0 1.75Zm5.125-.875a.875.875 0 1 1-1.75 0a.875.875 0 0 1 1.75 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}