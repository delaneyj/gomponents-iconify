package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOpenTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 5.5v6.604l1.5-2.6a3 3 0 0 1 2.599-1.5H15V7.5A1.5 1.5 0 0 0 13.5 6h-4a.5.5 0 0 1-.353-.146L7.44 4.146A.5.5 0 0 0 7.086 4H4.5A1.5 1.5 0 0 0 3 5.5Zm1.284 10.485c.071.01.144.015.22.015H13.9a2 2 0 0 0 1.732-1l2.163-3.745a1.5 1.5 0 0 0-1.3-2.25H7.1a2 2 0 0 0-1.732 1L3.205 13.75a1.501 1.501 0 0 0 1.079 2.235ZM2 14.46V5.5A2.5 2.5 0 0 1 4.5 3h2.586a1.5 1.5 0 0 1 1.06.44L9.708 5H13.5A2.5 2.5 0 0 1 16 7.5v.505h.496c1.925 0 3.128 2.083 2.166 3.75L16.498 15.5A3 3 0 0 1 13.9 17H4.5a2.457 2.457 0 0 1-1.624-.599A2.495 2.495 0 0 1 2 14.461Z"/>`),
		g.Group(children),
	)
}