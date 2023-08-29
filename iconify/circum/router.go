package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Router(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.44 10.94h-1.51v-6.4a.5.5 0 0 0-1 0v6.4H7.06V7a.5.5 0 0 0-1 0v3.94h-1.5a2.507 2.507 0 0 0-2.5 2.5v4a2.514 2.514 0 0 0 2.5 2.5h14.88a2.507 2.507 0 0 0 2.5-2.5v-4a2.5 2.5 0 0 0-2.5-2.5Zm1.5 6.5a1.5 1.5 0 0 1-1.5 1.5H4.56a1.511 1.511 0 0 1-1.5-1.5v-4a1.5 1.5 0 0 1 1.5-1.5h14.88a1.5 1.5 0 0 1 1.5 1.5Z"/><circle cx="4.75" cy="15.436" r=".75" fill="currentColor"/><circle cx="8.25" cy="15.436" r=".75" fill="currentColor"/><path fill="currentColor" d="M18.5 16.936h-5a1.5 1.5 0 1 1 0-3h5a1.5 1.5 0 0 1 0 3Zm-5-2a.5.5 0 1 0 0 1h5a.5.5 0 0 0 0-1Z"/>`),
		g.Group(children),
	)
}