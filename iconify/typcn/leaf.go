package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Leaf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 11c0-4.9-3.499-9.1-8.32-9.983L11.5.983l-.18.033a10.15 10.15 0 0 0-.82 19.779V22a1 1 0 1 0 2 0v-1.205A10.147 10.147 0 0 0 20 11zm-7.5 7.7v-2.993l4.354-4.354a.5.5 0 0 0-.707-.707L12.5 14.293v-3.586l2.354-2.354a.5.5 0 0 0-.707-.707L12.5 9.293V6a1 1 0 1 0-2 0v3.293L8.854 7.646a.5.5 0 0 0-.707.707l2.354 2.354v3.586l-3.646-3.646a.5.5 0 0 0-.707.707l4.354 4.354V18.7A8.144 8.144 0 0 1 11.5 3.019a8.145 8.145 0 0 1 1 15.681z"/>`),
		g.Group(children),
	)
}