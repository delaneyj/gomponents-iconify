package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InboxFull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 15V5H5v10h4c0 1.66 1.34 3 3 3s3-1.34 3-3h4m0-12c1.1 0 2 .9 2 2v14c0 1.1-.9 2-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14M7 13v-2h10v2H7m0-4V7h10v2H7Z"/>`),
		g.Group(children),
	)
}