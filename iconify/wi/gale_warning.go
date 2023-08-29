package wi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GaleWarning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 30 30"),
		g.Raw(`<path fill="currentColor" d="M10.67 24.6V7.45h1.03V24.6h-1.03zm1.73-2.16v-7.41l8.65 3.69l-8.65 3.72zm0-7.58V7.45l8.65 3.69l-8.65 3.72z"/>`),
		g.Group(children),
	)
}