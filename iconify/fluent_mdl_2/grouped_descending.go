package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GroupedDescending(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M384 1920v-128h1536v128H384zm0-1792h1536v128H512v1061l293-293l91 91l-448 447L0 1115l91-91l293 293V128zm640 640V640h896v128h-896zm0 384v-128h896v128h-896zm0 384v-128h896v128h-896z"/>`),
		g.Group(children),
	)
}