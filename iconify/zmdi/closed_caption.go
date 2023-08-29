package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClosedCaption(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M341 21q18 0 30.5 12.5T384 64v256q0 18-12.5 30.5T341 363H43q-18 0-30.5-12.5T0 320V64q0-18 12.5-30.5T43 21h298zM171 171v-22q0-8-6.5-14.5T149 128H85q-8 0-14.5 6.5T64 149v86q0 8 6.5 14.5T85 256h64q9 0 15.5-6.5T171 235v-22h-32v11H96v-64h43v11h32zm149 0v-22q0-8-6.5-14.5T299 128h-64q-9 0-15.5 6.5T213 149v86q0 8 6.5 14.5T235 256h64q8 0 14.5-6.5T320 235v-22h-32v11h-43v-64h43v11h32z"/>`),
		g.Group(children),
	)
}