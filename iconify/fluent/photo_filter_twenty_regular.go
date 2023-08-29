package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoFilterTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 8a6.003 6.003 0 0 0 4.257 5.743a6 6 0 1 0 7.486-7.486A6 6 0 0 0 2 8Zm6-5a5.001 5.001 0 0 1 4.597 3.03a6 6 0 0 0-6.567 6.567A5.001 5.001 0 0 1 8 3Zm9 9a5 5 0 0 1-9.597 1.97a6 6 0 0 0 6.567-6.567A5.001 5.001 0 0 1 17 12Z"/>`),
		g.Group(children),
	)
}