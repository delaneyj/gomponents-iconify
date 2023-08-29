package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tetragon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M3.575 13.388a1.962 1.962 0 0 1 0-2.776l7.037-7.037a1.962 1.962 0 0 1 2.776 0l7.037 7.037a1.963 1.963 0 0 1 0 2.776l-7.037 7.037a1.963 1.963 0 0 1-2.776 0l-7.037-7.037Z"/>`),
		g.Group(children),
	)
}