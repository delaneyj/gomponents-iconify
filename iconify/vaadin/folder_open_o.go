package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOpenO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14 6V4H7L6 2H2L1 4H0v11h14l2-9h-2zm.9 1l-1.6 7l-11.9-.1L3.7 7h11.2zM1 5h.6l1-2h2.6l1.2 2H13v1H3l-2 5.9V5z"/>`),
		g.Group(children),
	)
}