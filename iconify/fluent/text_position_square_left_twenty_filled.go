package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextPositionSquareLeftTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.5 2.75a.75.75 0 0 0 0 1.5h13a.75.75 0 0 0 0-1.5h-13Zm8.5 3a.75.75 0 0 0 0 1.5h4.5a.75.75 0 0 0 0-1.5H12Zm-.75 3.75a.75.75 0 0 1 .75-.75h4.5a.75.75 0 0 1 0 1.5H12a.75.75 0 0 1-.75-.75Zm.75 2.25a.75.75 0 0 0 0 1.5h4.5a.75.75 0 0 0 0-1.5H12ZM2.75 15.5a.75.75 0 0 1 .75-.75h13a.75.75 0 0 1 0 1.5h-13a.75.75 0 0 1-.75-.75Zm1.5-7a2.25 2.25 0 0 1 4.5 0v5a.75.75 0 0 0 1.5 0v-5a3.75 3.75 0 1 0-7.5 0v5a.75.75 0 0 0 1.5 0v-5Z"/>`),
		g.Group(children),
	)
}