package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M68.312 13.111H31.687a3.007 3.007 0 0 0-3.007 3.007v68.557c0 1.224.992 2.215 2.216 2.215a2.2 2.2 0 0 0 1.509-.602h.014l16.094-16.094h.035c.355-.414.876-.682 1.465-.682s1.109.268 1.465.682h.035l15.918 15.918a2.2 2.2 0 0 0 1.673.777a2.216 2.216 0 0 0 2.216-2.215V16.118a3.007 3.007 0 0 0-3.008-3.007z"/>`),
		g.Group(children),
	)
}