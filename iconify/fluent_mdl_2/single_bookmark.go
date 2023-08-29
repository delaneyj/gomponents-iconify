package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SingleBookmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1664 0v2048l-640-323l-640 323V0h1280zm-128 128H512v1712q129-65 256-130t256-129q129 64 256 129t256 130V128z"/>`),
		g.Group(children),
	)
}