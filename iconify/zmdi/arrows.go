package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Arrows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M192 176v-64h-64L235 5l106 107h-64v64h-85zm-21 21v86h-64v64L0 240l107-107v64h64zm298 43L363 347v-64h-64v-86h64v-64zm-192 64v64h64L235 475L128 368h64v-64h85z"/>`),
		g.Group(children),
	)
}