package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SearchReplace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M149 64q-38 0-67.5 24.5T45 149H2q8-54 49.5-91T149 21q62 0 106 44l44-44v128H171l54-54q-32-31-76-31zm121 195l103 104l-32 31l-103-103q-40 29-89 29q-62 0-105-44L0 320V192h128l-54 54q31 31 75 31q39 0 68-24t37-61h43q-5 37-27 67z"/>`),
		g.Group(children),
	)
}