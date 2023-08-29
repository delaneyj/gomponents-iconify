package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BinaryTreeTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 6a2 2 0 1 0-4 0a2 2 0 0 0 4 0zm-7 8a2 2 0 1 0-4 0a2 2 0 0 0 4 0zm14 0a2 2 0 1 0-4 0a2 2 0 0 0 4 0zm-7 4a2 2 0 1 0-4 0a2 2 0 0 0 4 0zM12 8v8m-5.684-3.504l4.368-4.992m7 4.992l-4.366-4.99"/>`),
		g.Group(children),
	)
}