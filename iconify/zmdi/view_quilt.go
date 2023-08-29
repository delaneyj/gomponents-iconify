package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewQuilt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M128 320V192h107v128H128zM0 320V43h107v277H0zm256 0V192h107v128H256zM128 43h235v128H128V43z"/>`),
		g.Group(children),
	)
}