package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M13.942 8.039a2.5 2.5 0 0 0-3.085-2.955a3 3 0 0 0-5.737.075A4 4 0 1 0 4 13h9.5a2.5 2.5 0 0 0 .442-4.961zM6.5 12L4 9.5l1-1L6.5 10L10 6.5l1 1L6.5 12z"/>`),
		g.Group(children),
	)
}