package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDoubleDownTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.854 8.646a.5.5 0 0 1 .001.707l-5.465 5.484a.55.55 0 0 1-.78 0L4.147 9.353a.5.5 0 1 1 .708-.706L10 13.812l5.147-5.165a.5.5 0 0 1 .707-.001Zm0-4a.5.5 0 0 1 .001.707l-5.465 5.484a.55.55 0 0 1-.78 0L4.147 5.353a.5.5 0 1 1 .708-.706L10 9.812l5.147-5.165a.5.5 0 0 1 .707-.001Z"/>`),
		g.Group(children),
	)
}