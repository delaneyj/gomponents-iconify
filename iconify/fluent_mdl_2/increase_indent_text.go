package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IncreaseIndentText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 640v128h-768V640h768zm-768-256h1024v128H1024V384zm768 768v128h-768v-128h768zm-768 384v-128h1024v128H1024zm0-512V896h1024v128H1024z"/>`),
		g.Group(children),
	)
}