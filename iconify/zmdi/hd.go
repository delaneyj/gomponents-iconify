package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M341 0q18 0 30.5 12.5T384 43v298q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0h298zM171 256V128h-32v53H96v-53H64v128h32v-43h43v43h32zm42-128v128h86q8 0 14.5-6.5T320 235v-86q0-8-6.5-14.5T299 128h-86zm32 96v-64h43v64h-43z"/>`),
		g.Group(children),
	)
}