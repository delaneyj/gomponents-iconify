package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignCenterVerticalSimpleThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 124h-36V48a12 12 0 0 0-12-12H96a12 12 0 0 0-12 12v76H48a4 4 0 0 0 0 8h36v76a12 12 0 0 0 12 12h64a12 12 0 0 0 12-12v-76h36a4 4 0 0 0 0-8Zm-44 84a4 4 0 0 1-4 4H96a4 4 0 0 1-4-4V48a4 4 0 0 1 4-4h64a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}