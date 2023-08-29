package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 1536v129h-385v383h-127v-383H383V512H0V385h383V0h129v385h1061l365-366l91 90l-366 366v1061h385zM512 512v933l933-933H512zm1024 1024V603l-933 933h933z"/>`),
		g.Group(children),
	)
}