package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockerMultiple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 2h18a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2m10 2v16h8V4h-8M3 4v16h8V4H3m2 9h2v4H5v-4m0-7h4v1.5H5V6m0 3h4v1.5H5V9m10 4h2v4h-2v-4m0-7h4v1.5h-4V6m0 3h4v1.5h-4V9Z"/>`),
		g.Group(children),
	)
}