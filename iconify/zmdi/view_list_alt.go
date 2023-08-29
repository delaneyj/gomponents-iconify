package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewListAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 213v-42h43v42H0zm0 86v-43h43v43H0zm0-171V85h43v43H0zm85 85v-42h299v42H85zm0 86v-43h299v43H85zm0-214h299v43H85V85z"/>`),
		g.Group(children),
	)
}