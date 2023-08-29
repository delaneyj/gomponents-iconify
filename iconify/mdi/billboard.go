package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Billboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 22h-2v-5h2M6 22H4v-5h2M23 4v9a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h18a2 2 0 0 1 2 2m-2 0H3v9h18m-1-7h-5v2h5m-2 1h-3v2h3m-4 0H4l2.73-3.64l2 2.73l.73-.54L8.2 7.82l1.71-2.27Z"/>`),
		g.Group(children),
	)
}