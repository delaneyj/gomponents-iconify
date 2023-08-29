package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewModule(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 171V43h107v128H0zm0 149V192h107v128H0zm128 0V192h107v128H128zm128 0V192h107v128H256zM128 171V43h107v128H128zM256 43h107v128H256V43z"/>`),
		g.Group(children),
	)
}