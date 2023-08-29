package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BinaryTree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 20a2 2 0 1 0-4 0a2 2 0 0 0 4 0zM16 4a2 2 0 1 0-4 0a2 2 0 0 0 4 0zm0 16a2 2 0 1 0-4 0a2 2 0 0 0 4 0zm-5-8a2 2 0 1 0-4 0a2 2 0 0 0 4 0zm10 0a2 2 0 1 0-4 0a2 2 0 0 0 4 0zM5.058 18.306l2.88-4.606m2.123-3.397l2.877-4.604m-2.873 8.006l2.876 4.6M15.063 5.7l2.881 4.61"/>`),
		g.Group(children),
	)
}