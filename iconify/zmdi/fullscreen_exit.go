package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FullscreenExit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 277v-42h107v106H64v-64H0zm64-170V43h43v106H0v-42h64zm128 234V235h107v42h-64v64h-43zm43-234h64v42H192V43h43v64z"/>`),
		g.Group(children),
	)
}