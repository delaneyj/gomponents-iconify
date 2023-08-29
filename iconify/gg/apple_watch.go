package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AppleWatch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16.497 5.03c0 .048-.001.096-.004.143A3.001 3.001 0 0 1 18.5 8.005v1h1v4h-1v3a3 3 0 0 1-2.005 2.83a3 3 0 0 1-2.995 3.17h-4a3 3 0 0 1-2.995-3.17a3.001 3.001 0 0 1-2.005-2.83v-8a3 3 0 0 1 2.003-2.83a3 3 0 0 1 3.01-3.18l4 .02a3 3 0 0 1 2.984 3.015Zm-8-.025h6a1 1 0 0 0-.995-.99l-4-.02a1 1 0 0 0-1.004.995v.015Zm7.208 2.021l-4.22-.021H7.5a1 1 0 0 0-1 1v8a1 1 0 0 0 .999 1H15.5a1 1 0 0 0 .999-1v-8a1 1 0 0 0-.795-.979ZM8.5 19.005a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1h-6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}