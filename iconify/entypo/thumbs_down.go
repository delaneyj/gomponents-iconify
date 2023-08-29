package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbsDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.352 12.638c.133.356-3.539 3.634-1.397 6.291c.501.621 2.201-2.975 4.615-4.602c1.331-.899 4.43-2.811 4.43-3.868V3.617C14 2.346 9.086 1 5.352 1C3.983 1 2 9.576 2 10.939c0 1.367 4.221 1.343 4.352 1.699zM15 12.543c.658 0 3-.4 3-3.123V4.572c0-2.721-2.342-3.021-3-3.021c-.657 0 1 .572 1 2.26v6.373c0 1.768-1.657 2.359-1 2.359z"/>`),
		g.Group(children),
	)
}