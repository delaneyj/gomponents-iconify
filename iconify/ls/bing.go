package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M340 180c188 0 340 101 340 226S528 633 340 633C153 633 1 532 0 407V75h98v172c62-42 147-67 242-67zm0 397c131 0 238-76 238-171c0-94-107-171-238-171s-238 77-238 171c0 95 107 171 238 171z"/>`),
		g.Group(children),
	)
}