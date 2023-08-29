package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartphoneSetting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M21 512v-43h43v43H21zm86 0v-43h42v43h-42zm85 0v-43h43v43h-43zM213 0q18 0 30.5 12.5T256 43v341q0 18-12.5 30.5T213 427H43q-18 0-30.5-12.5T0 384V43q0-18 12.5-30.5T43 0h170zm0 341V85H43v256h170z"/>`),
		g.Group(children),
	)
}