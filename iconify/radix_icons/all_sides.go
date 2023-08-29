package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AllSides(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.5.75L9.75 3h-4.5L7.5.75Zm0 13.5L9.75 12h-4.5l2.25 2.25Zm-4.5-9L.75 7.5L3 9.75v-4.5ZM14.25 7.5L12 5.25v4.5l2.25-2.25Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}