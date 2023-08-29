package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReportAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 1280h-384V384h384v896zM0 1024h384v640H0v-640zm1408 512h-128v128h-256V0h384v1536zM512 384h384v1280H512V384zm1536 1280v128h-256v256h-128v-256h-256v-128h256v-256h128v256h256z"/>`),
		g.Group(children),
	)
}