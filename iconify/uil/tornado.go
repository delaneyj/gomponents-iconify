package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tornado(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 21H8a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2Zm1-4H6a1 1 0 0 0 0 2h5a1 1 0 0 0 0-2Zm7-15a1 1 0 0 0-1-1H3a1 1 0 0 0 0 2h14a1 1 0 0 0 1-1Zm3 3H6a1 1 0 0 0 0 2h15a1 1 0 0 0 0-2Zm-2 4h-9a1 1 0 0 0 0 2h9a1 1 0 0 0 0-2Zm-5 4H8a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}