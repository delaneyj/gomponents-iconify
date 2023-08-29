package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.129 8.596l1.735 4.328l2.77 1.29L19 16.247V11.7l-8.871-3.104z" opacity=".7"/><path fill="currentColor" d="M14.634 14.214L9 17.457V3.4L5 2v17.76L9 22l10-5.753V11.7l-4.366 2.514z"/>`),
		g.Group(children),
	)
}