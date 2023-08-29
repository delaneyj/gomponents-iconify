package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquaredCross(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 2H4c-1.1 0-2 .9-2 2v12c0 1.1.9 2 2 2h12c1.1 0 2-.9 2-2V4c0-1.1-.9-2-2-2zm-2.939 12.789L10 11.729l-3.061 3.06l-1.729-1.728L8.271 10l-3.06-3.061L6.94 5.21L10 8.271l3.059-3.061l1.729 1.729L11.729 10l3.06 3.061l-1.728 1.728z"/>`),
		g.Group(children),
	)
}