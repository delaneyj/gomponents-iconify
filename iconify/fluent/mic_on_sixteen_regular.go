package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicOnSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M5.5 4.5a2.5 2.5 0 0 1 5 0V8a2.5 2.5 0 0 1-5 0V4.5zM8 3a1.5 1.5 0 0 0-1.5 1.5V8a1.5 1.5 0 1 0 3 0V4.5A1.5 1.5 0 0 0 8 3z" fill="currentColor"/><path d="M4 7.5a.5.5 0 0 1 .5.5a3.5 3.5 0 1 0 7 0a.5.5 0 0 1 1 0a4.5 4.5 0 0 1-4 4.473V13.5a.5.5 0 0 1-1 0v-1.027A4.5 4.5 0 0 1 3.5 8a.5.5 0 0 1 .5-.5z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}