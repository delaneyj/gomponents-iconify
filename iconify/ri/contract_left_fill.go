package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContractLeftFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.414 4.586V11H21v2h-5.586v6.414L8 12l7.414-7.414ZM4 19V5h2v14H4Z"/>`),
		g.Group(children),
	)
}