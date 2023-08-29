package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapExclamation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 20l-6-3l-6 3V7l6-3l6 3l6-3v8.5M9 4v13m6-10v13m4-4v3m0 3v.01"/>`),
		g.Group(children),
	)
}