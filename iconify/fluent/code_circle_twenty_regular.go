package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeCircleTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.958 6.2a.5.5 0 1 0-.916-.4l-3.5 8a.5.5 0 1 0 .916.4l3.5-8ZM6.854 7.646a.5.5 0 0 1 0 .708L5.207 10l1.647 1.646a.5.5 0 0 1-.708.708l-2-2a.5.5 0 0 1 0-.708l2-2a.5.5 0 0 1 .708 0Zm6.292.708a.5.5 0 0 1 .708-.708l2 2a.5.5 0 0 1 0 .708l-2 2a.5.5 0 0 1-.708-.708L14.793 10l-1.647-1.646ZM10 2a8 8 0 1 0 0 16a8 8 0 0 0 0-16Zm-7 8a7 7 0 1 1 14 0a7 7 0 0 1-14 0Z"/>`),
		g.Group(children),
	)
}