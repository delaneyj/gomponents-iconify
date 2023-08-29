package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7.847.14a.5.5 0 0 0-.694 0L2.867 4.283l-.004.003a6.237 6.237 0 0 0-1.747 3.23a6.12 6.12 0 0 0 .394 3.63a6.35 6.35 0 0 0 2.4 2.806A6.65 6.65 0 0 0 7.5 15a6.65 6.65 0 0 0 3.59-1.048a6.348 6.348 0 0 0 2.4-2.805a6.12 6.12 0 0 0 .394-3.63a6.238 6.238 0 0 0-1.747-3.23L7.847.14Z"/>`),
		g.Group(children),
	)
}