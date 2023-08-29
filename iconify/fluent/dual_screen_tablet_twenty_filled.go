package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DualScreenTabletTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2h-3v-2h.5a.5.5 0 0 0 0-1H13V9a3 3 0 0 0-3-3H4V5Zm6 12H7.5V7H10a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2ZM4 7h2.5v10H4a2 2 0 0 1-2-2V9a2 2 0 0 1 2-2Zm5.5 7a.5.5 0 0 0 0 1h.5a.5.5 0 0 0 0-1h-.5Zm-6 .5a.5.5 0 0 0 .5.5h.5a.5.5 0 0 0 0-1H4a.5.5 0 0 0-.5.5Z"/>`),
		g.Group(children),
	)
}