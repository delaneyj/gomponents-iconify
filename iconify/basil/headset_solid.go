package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadsetSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.285 8.344A2.751 2.751 0 0 0 2.25 11v3A2.75 2.75 0 0 0 5 16.75h2.5a.75.75 0 0 0 .75-.75V9a.75.75 0 0 0-.75-.75H5.815c.429-2.47 2.944-4.5 6.185-4.5c3.241 0 5.756 2.03 6.185 4.5H16.5a.75.75 0 0 0-.75.75v7c0 .414.336.75.75.75h1.663A3.251 3.251 0 0 1 15 19.25h-1.145a2 2 0 1 0 0 1.5H15c2.4 0 4.384-1.78 4.705-4.091A2.751 2.751 0 0 0 21.75 14v-3a2.751 2.751 0 0 0-2.035-2.656C19.332 4.84 15.926 2.25 12 2.25c-3.926 0-7.333 2.59-7.715 6.094Z"/>`),
		g.Group(children),
	)
}