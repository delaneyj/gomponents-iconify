package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBounceTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.75 6a.75.75 0 0 0-.75.75v6.5a.75.75 0 0 0 1.5 0V8.56l6.47 6.47a.75.75 0 0 0 1.06 0l6.754-6.754a.75.75 0 0 0-1.06-1.06L10.5 13.439L4.56 7.5h4.69a.75.75 0 0 0 0-1.5h-6.5Z"/>`),
		g.Group(children),
	)
}