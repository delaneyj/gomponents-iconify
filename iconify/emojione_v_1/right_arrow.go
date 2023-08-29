package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#1b75bb" d="M63.666 56.802a6.863 6.863 0 0 1-6.862 6.868H6.86A6.864 6.864 0 0 1 0 56.802V6.864A6.863 6.863 0 0 1 6.86 0h49.943a6.862 6.862 0 0 1 6.862 6.864v49.938z"/><path fill="#fff" d="m35.719 15.915l18.428 15.858l-18.25 15.936c-2.56-.189-4.57-1.893-4.582-3.979l-.018-3.541l-16.625.09a5.074 5.074 0 0 1-5.096-5.04l-.031-6.572a5.067 5.067 0 0 1 5.04-5.092l16.622-.086l-.02-3.54c-.009-2.087 1.974-3.809 4.53-4.03"/>`),
		g.Group(children),
	)
}