package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomOutTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 8.5a5.5 5.5 0 1 0-2.177 4.383l4.147 4.147l.084.073a.75.75 0 0 0 .976-1.133l-4.147-4.147A5.476 5.476 0 0 0 14 8.5Zm-3-.75a.75.75 0 0 1 0 1.5H6a.75.75 0 0 1 0-1.5h5Z"/>`),
		g.Group(children),
	)
}