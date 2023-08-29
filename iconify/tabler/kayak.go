package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kayak(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6.414 6.414a2 2 0 0 0 0-2.828L5 2.172L2.172 5l1.414 1.414a2 2 0 0 0 2.828 0zm11.172 11.172a2 2 0 0 0 0 2.828L19 21.828L21.828 19l-1.414-1.414a2 2 0 0 0-2.828 0zM6.5 6.5l11 11m4.5-15C12.017 5.101 4.373 10.452 2 22c9.983-2.601 17.627-7.952 20-19.5zm-15.5 10l5 5m1-11l5 5"/>`),
		g.Group(children),
	)
}