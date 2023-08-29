package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddBookmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1024 1581l64 32q32 16 64 33v143l-128-64l-640 323V0h1280v1280h-128V128H512v1712q129-65 256-130t256-129zm896 83v128h-256v256h-128v-256h-256v-128h256v-256h128v256h256z"/>`),
		g.Group(children),
	)
}