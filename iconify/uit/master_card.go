package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MasterCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.25 5.204a6.67 6.67 0 0 0-3.25.843a6.67 6.67 0 0 0-3.25-.843a6.796 6.796 0 0 0 0 13.592a6.67 6.67 0 0 0 3.25-.843a6.67 6.67 0 0 0 3.25.843a6.796 6.796 0 0 0 0-13.592zm-6.5 12.592a5.796 5.796 0 0 1 0-11.592c.792 0 1.575.166 2.298.487a6.805 6.805 0 0 0 0 10.618a5.676 5.676 0 0 1-2.298.487zm3.25-1.02A5.805 5.805 0 0 1 9.5 12A5.805 5.805 0 0 1 12 7.223a5.813 5.813 0 0 1 0 9.554zm3.25 1.02a5.676 5.676 0 0 1-2.298-.487a6.805 6.805 0 0 0 0-10.618a5.676 5.676 0 0 1 2.298-.487a5.796 5.796 0 0 1 0 11.592z"/>`),
		g.Group(children),
	)
}