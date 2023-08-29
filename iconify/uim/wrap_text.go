package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WrapText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 7.167H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm-12 10H3a1 1 0 1 1 0-2h6a1 1 0 0 1 0 2z" opacity=".5"/><path fill="currentColor" d="M18.5 17.167H15a1 1 0 0 1 0-2h3.5a1.5 1.5 0 0 0 0-3H3a1 1 0 0 1 0-2h15.5a3.5 3.5 0 1 1 0 7Z"/><path fill="currentColor" d="M14.999 18.833a.995.995 0 0 1-.639-.231l-2-1.666a1 1 0 0 1 0-1.538l2-1.667a1 1 0 0 1 1.28 1.538l-1.078.898l1.078.897a1 1 0 0 1-.641 1.769Z"/>`),
		g.Group(children),
	)
}