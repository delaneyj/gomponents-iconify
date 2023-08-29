package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Http(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M75 171v-43h32v128H75v-53H32v53H0V128h32v43h43zm53-11v-32h96v32h-32v96h-32v-96h-32zm117 0v-32h96v32h-32v96h-32v-96h-32zm192-32q13 0 22.5 9.5T469 160v21q0 13-9.5 22.5T437 213h-42v43h-32V128h74zm0 53v-21h-42v21h42z"/>`),
		g.Group(children),
	)
}