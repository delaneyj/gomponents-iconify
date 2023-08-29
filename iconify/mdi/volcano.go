package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Volcano(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 8h-7l-2 5H6l-4 9h20L18 8m-5-7h2v4h-2V1m3.12 4.47l2.83-2.83l1.41 1.41l-2.82 2.83l-1.42-1.41M7.64 4.05l1.41-1.41l2.83 2.82l-1.41 1.42l-2.83-2.83Z"/>`),
		g.Group(children),
	)
}