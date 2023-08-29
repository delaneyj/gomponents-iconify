package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BugFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M168 92a12 12 0 1 1-12-12a12 12 0 0 1 12 12Zm-68-12a12 12 0 1 0 12 12a12 12 0 0 0-12-12Zm116 64a87.76 87.76 0 0 1-3 23l22.24 9.72A8 8 0 0 1 232 192a7.89 7.89 0 0 1-3.2-.67L207.38 182a88 88 0 0 1-158.76 0l-21.42 9.33a7.89 7.89 0 0 1-3.2.67a8 8 0 0 1-3.2-15.33L43 167a87.76 87.76 0 0 1-3-23v-8H16a8 8 0 0 1 0-16h24v-8a87.76 87.76 0 0 1 3-23l-22.2-9.67a8 8 0 1 1 6.4-14.66L48.62 74a88 88 0 0 1 158.76 0l21.42-9.36a8 8 0 0 1 6.4 14.66L213 89.05a87.76 87.76 0 0 1 3 23v8h24a8 8 0 0 1 0 16h-24Zm-80 0a8 8 0 0 0-16 0v64a8 8 0 0 0 16 0Zm64-32a72 72 0 0 0-144 0v8h144Z"/>`),
		g.Group(children),
	)
}