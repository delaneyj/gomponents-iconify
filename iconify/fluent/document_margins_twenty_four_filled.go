package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentMarginsTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.5 2h-7v4.25a.75.75 0 0 1-1.5 0V2h-.75A2.25 2.25 0 0 0 4 4.25v15.5A2.25 2.25 0 0 0 6.25 22H7v-4.25a.75.75 0 0 1 1.5 0V22h7v-4.25a.75.75 0 0 1 1.5 0V22h.75A2.25 2.25 0 0 0 20 19.75V4.25A2.25 2.25 0 0 0 17.75 2H17v4.25a.75.75 0 0 1-1.5 0V2ZM7.75 8.5a.75.75 0 0 1 .75.75v5.5a.75.75 0 0 1-1.5 0v-5.5a.75.75 0 0 1 .75-.75Zm9.25.75v5.5a.75.75 0 0 1-1.5 0v-5.5a.75.75 0 0 1 1.5 0Z"/>`),
		g.Group(children),
	)
}