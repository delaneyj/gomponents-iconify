package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Subway(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M10 4C7.25 4 5 6.25 5 9v12a6.012 6.012 0 0 0 3.531 5.469L6 29h2.344l2.031-2.031c.2.02.418.031.625.031h10c.207 0 .426-.012.625-.031L23.656 29H26l-2.531-2.531A6.012 6.012 0 0 0 27 21V9c0-2.75-2.25-5-5-5zm0 2h12c1.32 0 2.438.828 2.844 2H7.156A2.992 2.992 0 0 1 10 6zm-3 4h8v6H7zm10 0h8v6h-8zM7 18h18v3c0 2.219-1.781 4-4 4H11c-2.219 0-4-1.781-4-4zm3.5 2a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3zm11 0a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3z"/>`),
		g.Group(children),
	)
}