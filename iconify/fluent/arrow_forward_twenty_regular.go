package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowForwardTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m16.293 9l-3.39 3.39a.5.5 0 0 0 .639.764l.069-.057l4.243-4.243a.5.5 0 0 0 .057-.638l-.057-.07l-4.243-4.242a.5.5 0 0 0-.765.638l.058.07L16.293 8H10a7.5 7.5 0 0 0-7.496 7.258L2.5 15.5a.5.5 0 0 0 1 0a6.5 6.5 0 0 1 6.267-6.496L10 9h6.293Z"/>`),
		g.Group(children),
	)
}