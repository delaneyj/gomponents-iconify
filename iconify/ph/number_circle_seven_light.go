package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleSevenLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 26a102 102 0 1 0 102 102A102.12 102.12 0 0 0 128 26Zm0 192a90 90 0 1 1 90-90a90.1 90.1 0 0 1-90 90Zm28.91-133.44a6 6 0 0 1 .73 5.49l-32 88A6 6 0 0 1 120 182a6.15 6.15 0 0 1-2-.36a6 6 0 0 1-3.59-7.69L143.43 94H104a6 6 0 0 1 0-12h48a6 6 0 0 1 4.91 2.56Z"/>`),
		g.Group(children),
	)
}