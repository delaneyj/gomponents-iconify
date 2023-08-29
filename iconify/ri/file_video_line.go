package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileVideoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 4v4h4v12H5V4h10ZM3.998 2A.995.995 0 0 0 3 2.992v18.016a1 1 0 0 0 .993.992h16.014A1 1 0 0 0 21 20.992V7l-5-5H3.998Zm11.003 9.667l-4.88-3.253a.4.4 0 0 0-.621.333v6.505a.4.4 0 0 0 .622.333L15 12.333a.401.401 0 0 0 0-.666Z"/>`),
		g.Group(children),
	)
}