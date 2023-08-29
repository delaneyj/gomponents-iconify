package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M256 0q18 0 30.5 12.5T299 43v341l-150-64L0 384V43q0-18 12.5-30.5T43 0h213z"/>`),
		g.Group(children),
	)
}