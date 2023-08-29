package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableComputed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1062 1280H640v256h293l-65 128H0V128h1664v512h-128V512h-384v256h168l-65 128h-103v205l-90 179zm-38-128V896H640v256h384zm0-384V512H640v256h384zm-512 768v-256H128v256h384zm0-384V896H128v256h384zm0-384V512H128v256h384zM128 384h1408V256H128v128zm1651 768h269l-672 896h-264l256-512h-256l387-768h518l-238 384zm-444 738l457-610h-243l238-384h-209l-258 512h256l-241 482z"/>`),
		g.Group(children),
	)
}