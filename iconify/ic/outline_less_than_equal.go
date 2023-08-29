package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutlineLessThanEqual(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17.5 15.5L9.25 10l8.25-5.5l-1-1.5L6 10l10.5 7z"/><path fill="currentColor" d="M18 20.998H6v-2h12z"/>`),
		g.Group(children),
	)
}