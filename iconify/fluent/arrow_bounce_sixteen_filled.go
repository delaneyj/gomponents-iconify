package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBounceSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.002 4.75a.75.75 0 0 0-.75-.75h-5.5a.75.75 0 0 0-.75.75v5.5a.75.75 0 0 0 1.5 0V6.561l5.718 5.72a.75.75 0 0 0 1.06 0l5.5-5.5a.75.75 0 1 0-1.06-1.061l-4.97 4.97L3.562 5.5h3.69a.75.75 0 0 0 .75-.75Z"/>`),
		g.Group(children),
	)
}