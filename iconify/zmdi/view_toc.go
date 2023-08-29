package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewToc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 128V85h299v43H0zm0 85v-42h299v42H0zm0 86v-43h299v43H0zm341 0v-43h43v43h-43zm0-214h43v43h-43V85zm0 128v-42h43v42h-43z"/>`),
		g.Group(children),
	)
}