package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonInfoTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11 17.5a6.47 6.47 0 0 1 1.022-3.5h-6.77a2.249 2.249 0 0 0-2.248 2.249v.92c0 .572.178 1.13.51 1.596C5.056 20.929 7.58 22 11 22c.6 0 1.173-.033 1.717-.099A6.477 6.477 0 0 1 11 17.5Zm0-15.495a5 5 0 1 1 0 10a5 5 0 0 1 0-10ZM23 17.5a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0Zm-9.5 0a4 4 0 1 0 8 0a4 4 0 0 0-8 0Zm5-2a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-.25 2.5a.75.75 0 0 0-1.5 0v1.75a.75.75 0 0 0 1.5 0V18Z"/>`),
		g.Group(children),
	)
}