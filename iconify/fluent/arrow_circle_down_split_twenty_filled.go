package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleDownSplitTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18 10a8 8 0 1 1-16 0a8 8 0 0 1 16 0Zm-9.207 2.414a.5.5 0 1 0-.707-.707l-.586.586V9h5v3.293l-.586-.586a.5.5 0 0 0-.707.707l1.44 1.44a.5.5 0 0 0 .707 0l1.439-1.44a.5.5 0 0 0-.707-.707l-.586.586V8.5A.5.5 0 0 0 13 8h-2.5V5.5a.5.5 0 0 0-1 0V8H7a.5.5 0 0 0-.5.5v3.793l-.586-.586a.5.5 0 0 0-.707.707l1.44 1.44a.5.5 0 0 0 .707 0l1.439-1.44Z"/>`),
		g.Group(children),
	)
}