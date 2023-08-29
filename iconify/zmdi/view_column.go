package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewColumn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M128 320V43h107v277H128zM0 320V43h107v277H0zM256 43h107v277H256V43z"/>`),
		g.Group(children),
	)
}