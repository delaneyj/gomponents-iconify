package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Undo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 1a7.979 7.979 0 0 0-5.657 2.343L0 1v6h6L3.757 4.757a6 6 0 1 1 8.211 8.743l1.323 1.5A8 8 0 0 0 8 1z"/>`),
		g.Group(children),
	)
}