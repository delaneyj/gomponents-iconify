package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowEnterUpTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.864 2.157a.5.5 0 0 0-.728 0l-4 4.25a.5.5 0 0 0 .728.686L8 3.76V15a3 3 0 0 0 3 3h4.5a.5.5 0 0 0 0-1H11a2 2 0 0 1-2-2V3.76l3.136 3.333a.5.5 0 0 0 .728-.686l-4-4.25Z"/>`),
		g.Group(children),
	)
}