package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Medal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.38 5.76a1 1 0 0 0-.47-.61l-5.2-3a1 1 0 0 0-1.37.36L12 6.57L9.66 2.51a1 1 0 0 0-1.37-.36l-5.2 3a1 1 0 0 0-.47.61a1 1 0 0 0 .1.75l4 6.83A5.91 5.91 0 0 0 6 16a6 6 0 1 0 11.34-2.72l3.9-6.76a1 1 0 0 0 .14-.76ZM5 6.38l3.46-2L11.68 10A5.94 5.94 0 0 0 8 11.58ZM12 20a4 4 0 0 1-4-4a4 4 0 0 1 4-4a4 4 0 1 1 0 8Zm4-8.45a5.9 5.9 0 0 0-1.86-1.15l-.98-1.83l2.42-4.19l3.46 2Z"/>`),
		g.Group(children),
	)
}