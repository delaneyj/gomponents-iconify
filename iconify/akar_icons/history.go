package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func History(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.266 16.06a8.923 8.923 0 0 0 3.915 3.978a8.706 8.706 0 0 0 5.471.832a8.795 8.795 0 0 0 4.887-2.64a9.067 9.067 0 0 0 2.388-5.079a9.135 9.135 0 0 0-1.044-5.53a8.903 8.903 0 0 0-4.069-3.815a8.7 8.7 0 0 0-5.5-.608c-1.85.401-3.366 1.313-4.62 2.755c-.151.16-.735.806-1.22 1.781M7.5 8l-3.609.72L3 5m9 4v4l3 2"/>`),
		g.Group(children),
	)
}