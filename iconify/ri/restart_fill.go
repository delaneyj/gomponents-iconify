package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RestartFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Zm4.82-4.924a7 7 0 1 0-1.853 1.266l-.974-1.755A5 5 0 1 1 17 12h-3l2.82 5.076Z"/>`),
		g.Group(children),
	)
}