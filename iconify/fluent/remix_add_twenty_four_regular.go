package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemixAddTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 2.75A.75.75 0 0 1 2.75 2H12c5.523 0 10 4.477 10 10a9.976 9.976 0 0 1-3.385 7.5h-2.611A8.503 8.503 0 0 0 12 3.5H2.75A.75.75 0 0 1 2 2.75Zm8.362 19.116c.533.088 1.08.134 1.638.134h9.25a.75.75 0 0 0 0-1.5H12A8.496 8.496 0 0 1 3.5 12a8.499 8.499 0 0 1 4.496-7.5h-2.61A9.976 9.976 0 0 0 2 12c0 4.965 3.618 9.085 8.362 9.866ZM12 8a.75.75 0 0 1 .75.75v2.5h2.5a.75.75 0 0 1 0 1.5h-2.5v2.5a.75.75 0 0 1-1.5 0v-2.5h-2.5a.75.75 0 0 1 0-1.5h2.5v-2.5A.75.75 0 0 1 12 8Z"/>`),
		g.Group(children),
	)
}