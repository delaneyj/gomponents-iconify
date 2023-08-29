package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkTwoSd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.5 11V6.5a2 2 0 0 1 2-2h27a2 2 0 0 1 2 2v35a2 2 0 0 1-2 2h-18l-11-11V18"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.84 14.5L12 7.5V11H5v7h7v3.5l8.84-7zm13.05 18.85h3v5h-3zm-12.29 0h3v5h-3zm6.15 0h3v5h-3zm-9.29 0v5l-4.97-4.97l4.97-.03z"/>`),
		g.Group(children),
	)
}