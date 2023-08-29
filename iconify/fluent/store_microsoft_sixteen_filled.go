package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StoreMicrosoftSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 3.5V5H2.5a.5.5 0 0 0-.5.5V12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V5.5a.5.5 0 0 0-.5-.5H11V3.5A1.5 1.5 0 0 0 9.5 2h-3A1.5 1.5 0 0 0 5 3.5ZM6.5 3h3a.5.5 0 0 1 .5.5V5H6V3.5a.5.5 0 0 1 .5-.5Zm-1 6V7h2v2h-2Zm0 3v-2h2v2h-2Zm5-3h-2V7h2v2Zm-2 3v-2h2v2h-2Z"/>`),
		g.Group(children),
	)
}