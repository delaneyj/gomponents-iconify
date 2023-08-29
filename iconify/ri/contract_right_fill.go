package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContractRightFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.586 4.586L16 12l-7.414 7.414V13H3v-2h5.586V4.586ZM18 19V5h2v14h-2Z"/>`),
		g.Group(children),
	)
}