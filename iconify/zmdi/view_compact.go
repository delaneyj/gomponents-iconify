package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewCompact(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 341V192h128v149H0zm149 0V192h256v149H149zM0 43h405v128H0V43z"/>`),
		g.Group(children),
	)
}