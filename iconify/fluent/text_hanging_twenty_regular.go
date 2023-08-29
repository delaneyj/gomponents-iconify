package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextHangingTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16.5 4a.5.5 0 0 1 0 1h-13a.5.5 0 0 1 0-1h13Zm-4 10a.5.5 0 0 1 0 1h-9a.5.5 0 0 1 0-1h9Zm.5-4.5a.5.5 0 0 0-.5-.5h-9a.5.5 0 0 0 0 1h9a.5.5 0 0 0 .5-.5Zm2.146 2.146l1.5-1.5a.5.5 0 0 1 .708.708L16.207 12l1.147 1.146a.5.5 0 0 1-.708.708l-1.5-1.5a.5.5 0 0 1 0-.708Z"/>`),
		g.Group(children),
	)
}