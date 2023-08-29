package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Download(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M40.5 38.5v-1c0-1.48-.43-2-2-2h-34c-1.48 0-2 .49-2 2v1c0 1.55.52 2 2 2h34c1.51 0 2-.48 2-2zm-15-36h-8c-2.5 0-3 .47-3 3v11h-10l17 17l17-17h-10v-11c0-2.53-.529-3-3-3z"/>`),
		g.Group(children),
	)
}