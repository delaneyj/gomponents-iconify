package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BackwardStartCircleLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle r="10" transform="matrix(-1 0 0 1 12 12)"/><path d="M17 15V9l-5 3l5 3zM7 12l5 3V9l-5 3zm0 0V9m0 3v3"/></g>`),
		g.Group(children),
	)
}