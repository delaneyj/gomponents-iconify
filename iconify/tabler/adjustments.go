package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Adjustments(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 10a2 2 0 1 0 4 0a2 2 0 0 0-4 0m2-6v4m0 4v8m4-4a2 2 0 1 0 4 0a2 2 0 0 0-4 0m2-12v10m0 4v2m4-13a2 2 0 1 0 4 0a2 2 0 0 0-4 0m2-3v1m0 4v11"/>`),
		g.Group(children),
	)
}