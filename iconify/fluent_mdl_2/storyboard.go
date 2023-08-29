package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Storyboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 128v1664H0V128h2048zM128 256v256h1792V256H128zm0 1408h1024V640H128v1024zm1792 0V640h-640v1024h640zm-512-896h384v128h-384V768zm0 256h384v128h-384v-128zm0 256h384v128h-384v-128zm-384-512v768H256V768h768zM896 896H384v512h512V896z"/>`),
		g.Group(children),
	)
}