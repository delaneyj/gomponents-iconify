package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonQuestionMarkSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.5 9a1.5 1.5 0 0 1 1.5 1.5v.5c0 1.971-1.86 4-5 4c-3.14 0-5-2.029-5-4v-.5A1.5 1.5 0 0 1 2.5 9h7ZM13 7a.75.75 0 1 1 0 1.5a.75.75 0 0 1 0-1.5ZM6 2.5A2.75 2.75 0 1 1 6 8a2.75 2.75 0 0 1 0-5.5ZM13 1a2 2 0 0 1 2 2c0 .73-.212 1.14-.754 1.708l-.264.27c-.377.393-.482.605-.482 1.022a.5.5 0 0 1-1 0c0-.73.212-1.14.754-1.708l.264-.27C13.895 3.63 14 3.418 14 3a1 1 0 1 0-2 0a.5.5 0 0 1-1 0a2 2 0 0 1 2-2Z"/>`),
		g.Group(children),
	)
}