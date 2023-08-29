package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimerOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 3.018a.75.75 0 1 1 0-1.5h3.536a.75.75 0 0 1 0 1.5H10Zm-3.47.452a.75.75 0 0 1 0 1.06l-2.5 2.5a.75.75 0 0 1-1.06-1.06l2.5-2.5a.75.75 0 0 1 1.06 0ZM12 5.75A7.25 7.25 0 1 0 19.25 13a.75.75 0 1 1 1.5 0A8.75 8.75 0 1 1 12 4.25a.75.75 0 0 1 0 1.5Z"/><path fill="currentColor" d="M17.188 8.364a.75.75 0 0 0-1.052-1.052l-3.17 2.465l-2.072 1.48a1.684 1.684 0 1 0 2.35 2.349l1.479-2.072l2.465-3.17Z"/>`),
		g.Group(children),
	)
}