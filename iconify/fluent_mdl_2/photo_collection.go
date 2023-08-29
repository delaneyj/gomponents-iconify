package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoCollection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1472 640q-26 0-45-19t-19-45q0-26 19-45t45-19q26 0 45 19t19 45q0 26-19 45t-45 19zm576-128v1280H256v-256H0V256h1792v256h256zm-475 896l-357-358l-166 166l193 192h330zM128 384v549l320-319l512 512l256-256l448 447V384H128zm933 1024L448 794l-320 321v293h933zm859-768h-128v896H384v128h1536V640z"/>`),
		g.Group(children),
	)
}