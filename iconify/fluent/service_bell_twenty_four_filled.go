package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ServiceBellTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 4a2 2 0 1 1 4 0v1.24a8.24 8.24 0 0 1 6.25 8.008a.75.75 0 0 1-.75.75h-15a.75.75 0 0 1-.75-.75c0-3.86 2.668-7.098 6.25-7.999V4Zm3.25 3.5a.75.75 0 0 0 0 1.5c1.253 0 2.143.618 2.886 1.68a.75.75 0 0 0 1.229-.86c-.93-1.33-2.229-2.32-4.115-2.32ZM2 17a2 2 0 0 1 2-2h16a2 2 0 1 1 0 4H4a2 2 0 0 1-2-2Z"/>`),
		g.Group(children),
	)
}