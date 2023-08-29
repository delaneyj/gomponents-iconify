package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 0h384v384H0V0zm171 341V213H43v128h128zm0-170V43H43v128h128zm170 170V213H213v128h128zm0-170V43H213v128h128z"/>`),
		g.Group(children),
	)
}