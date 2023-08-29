package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarHalf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M427 160L310 261l35 150l-132-80l-132 80l35-150L0 160l153-13L213 5l60 142zM213 291l81 49l-22-91l71-62l-93-8l-37-86v198z"/>`),
		g.Group(children),
	)
}