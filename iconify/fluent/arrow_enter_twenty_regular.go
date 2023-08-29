package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowEnterTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.5 17a.5.5 0 0 1 0-1H14a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2H7.5a.5.5 0 0 1 0-1H14a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3H7.5Zm.146-10.854a.5.5 0 0 1 .708 0l3.5 3.5a.5.5 0 0 1 0 .708l-3.5 3.5a.5.5 0 0 1-.708-.708l2.647-2.646H1.5a.5.5 0 0 1 0-1h8.793L7.646 6.854a.5.5 0 0 1 0-.708Z"/>`),
		g.Group(children),
	)
}