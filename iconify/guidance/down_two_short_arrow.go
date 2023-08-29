package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownTwoShortArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M1.5 12c1.11 0 2.771 1.1 4.166 2.212c1.796 1.432 3.365 3.141 4.563 5.102c.897 1.469 1.758 3.253 1.758 4.686M22.5 12c-1.11 0-2.771 1.1-4.166 2.212c-1.796 1.432-3.365 3.141-4.563 5.102c-.897 1.469-1.758 3.253-1.758 4.686M12 24V0"/>`),
		g.Group(children),
	)
}