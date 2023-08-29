package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hq(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M341 21q18 0 30.5 12.5T384 64v256q0 18-12.5 30.5T341 363H43q-18 0-30.5-12.5T0 320V64q0-18 12.5-30.5T43 21h298zM171 256V128h-32v53H96v-53H64v128h32v-43h43v43h32zm149-21v-86q0-8-6.5-14.5T299 128h-64q-9 0-15.5 6.5T213 149v86q0 8 6.5 14.5T235 256h16v32h32v-32h16q8 0 14.5-6.5T320 235zm-75-11v-64h43v64h-43z"/>`),
		g.Group(children),
	)
}