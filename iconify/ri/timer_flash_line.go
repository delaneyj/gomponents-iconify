package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimerFlashLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.382 5.968A8.962 8.962 0 0 1 12 4c2.125 0 4.078.736 5.618 1.968l1.453-1.453l1.414 1.414l-1.453 1.453A9 9 0 1 1 3 13c0-2.125.736-4.078 1.968-5.618L3.515 5.93l1.414-1.414l1.453 1.453ZM12 20a7 7 0 1 0 0-14a7 7 0 0 0 0 14Zm1-8h3l-5 6.5V14H8l5-6.505V12ZM8 1h8v2H8V1Z"/>`),
		g.Group(children),
	)
}