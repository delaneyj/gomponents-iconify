package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Swap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M85 171v64h150v42H85v64L0 256zm299-43l-85 85v-64H149v-42h150V43z"/>`),
		g.Group(children),
	)
}