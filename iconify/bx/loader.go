package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Loader(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 11h5v2H2zm15 0h5v2h-5zm-6 6h2v5h-2zm0-15h2v5h-2zM4.222 5.636l1.414-1.414l3.536 3.536l-1.414 1.414zm15.556 12.728l-1.414 1.414l-3.536-3.536l1.414-1.414zm-12.02-3.536l1.414 1.414l-3.536 3.536l-1.414-1.414zm7.07-7.071l3.536-3.535l1.414 1.415l-3.536 3.535z"/>`),
		g.Group(children),
	)
}