package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkiResorts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1472 640l574 1152H0L768 256l447 897l257-513zm0 287l-185 369l47 95l111-111h203l-176-353zM898 803L768 543L638 803l130 130l130-130zm-691 861h1121L958 924l-190 191l-191-191l-370 740zm1263 0h369l-127-256h-213l-104 104l75 152z"/>`),
		g.Group(children),
	)
}