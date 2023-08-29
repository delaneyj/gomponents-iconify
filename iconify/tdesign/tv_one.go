package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TvOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 3h24v15h-4.046l2.698 2.063l-1.215 1.589L16.66 18H7.335l-4.771 3.652l-1.216-1.588L4.044 18H0V3Zm2 2v11h20V5H2Z"/>`),
		g.Group(children),
	)
}