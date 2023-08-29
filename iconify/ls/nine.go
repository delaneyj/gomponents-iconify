package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="m82 743l186-292c-13 2-27 4-41 4C101 455 0 353 0 228C0 102 101 0 227 0c105 0 194 73 220 171c4 17 8 38 8 57c0 43-13 82-33 116l-7 11l-273 427zm266-417l19-29c11-22 16-44 16-69c0-86-70-157-156-157S71 142 71 228s70 155 156 155c49 0 92-22 121-57z"/>`),
		g.Group(children),
	)
}