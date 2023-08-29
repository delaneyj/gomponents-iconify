package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipForwardMiniFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.788 17.444A.5.5 0 0 1 7 17.035V6.965a.5.5 0 0 1 .788-.409l7.133 5.035a.499.499 0 0 1 0 .818l-7.133 5.034ZM16 7a1 1 0 1 1 2 0v10a1 1 0 1 1-2 0V7Z"/>`),
		g.Group(children),
	)
}