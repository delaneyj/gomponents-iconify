package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionForkFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" d="M8 14V9.414l7 7V26h2v-9.586l7-7V14h2V6h-8v2h4.586L16 14.586L9.414 8H14V6H6v8h2z"/><path fill="currentColor" d="M28 2H4a2 2 0 0 0-2 2v24a2 2 0 0 0 2 2h24a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2Zm-2 12h-2V9.414l-7 7V26h-2v-9.586l-7-7V14H6V6h8v2H9.414L16 14.586L22.586 8H18V6h8Z"/>`),
		g.Group(children),
	)
}