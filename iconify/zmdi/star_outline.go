package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M427 157L310 258l35 150l-132-80l-132 80l35-150L0 157l153-13L213 3l60 141zM213 289l81 48l-22-91l71-62l-93-8l-37-86l-36 86l-93 8l70 62l-21 91z"/>`),
		g.Group(children),
	)
}