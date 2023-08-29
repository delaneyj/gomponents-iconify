package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContractLeftLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.071 4.929l1.415 1.414L11.829 11H21v2h-9.172l4.657 4.657l-1.415 1.414L8.001 12l7.07-7.071Zm-11.07 14.07V5h2v14H4Z"/>`),
		g.Group(children),
	)
}