package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContractRightLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.171 11L7.515 6.343l1.414-1.414l7.07 7.07l-7.07 7.072l-1.414-1.414L12.17 13H3v-2h9.171ZM18 19V5h2v14h-2Z"/>`),
		g.Group(children),
	)
}