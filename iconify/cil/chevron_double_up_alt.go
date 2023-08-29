package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDoubleUpAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M20.095 20.405L12 12.31l-8.095 8.095l-1.061-1.061l9.155-9.155l9.155 9.155l-1.061 1.061z" fill="currentColor"/><path d="M20.095 12.905L12 4.81l-8.095 8.095l-1.061-1.061l9.155-9.155l9.155 9.155l-1.061 1.061z" fill="currentColor"/>`),
		g.Group(children),
	)
}