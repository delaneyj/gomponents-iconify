package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowExpandHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 11h6V8l4 4l-4 4v-3H9v3l-4-4l4-4v3m-7 9V4h2v16H2m18 0V4h2v16h-2Z"/>`),
		g.Group(children),
	)
}