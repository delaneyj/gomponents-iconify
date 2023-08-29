package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurveArray(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M5.164 17a17.47 17.47 0 0 1 1.132-3M11.5 7.794A16.838 16.838 0 0 1 14 6.296M4.5 22a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5Z"/><path d="M9.5 12a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5Zm10-5a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5Z"/></g>`),
		g.Group(children),
	)
}