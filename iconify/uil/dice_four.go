package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiceFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 14a1 1 0 1 0 1 1a1 1 0 0 0-1-1ZM9 8a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm0 6a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm6-6a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm2-6H7a5 5 0 0 0-5 5v10a5 5 0 0 0 5 5h10a5 5 0 0 0 5-5V7a5 5 0 0 0-5-5Zm3 15a3 3 0 0 1-3 3H7a3 3 0 0 1-3-3V7a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3Z"/>`),
		g.Group(children),
	)
}