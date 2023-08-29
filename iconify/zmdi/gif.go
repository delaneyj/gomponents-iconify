package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gif(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M139 128h32v128h-32V128zm-54 0q10 0 16 6.5t6 14.5v11H32v64h43v-32h32v43q0 8-6 14.5T85 256H21q-9 0-15-6.5T0 235v-86q0-8 6-14.5t15-6.5h64zm214 32h-64v21h42v32h-42v43h-32V128h96v32z"/>`),
		g.Group(children),
	)
}