package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarbonAccounting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M29 26h-6v-4a2.002 2.002 0 0 1 2-2h2v-2h-4v-2h4a2.002 2.002 0 0 1 2 2v2a2.002 2.002 0 0 1-2 2h-2v2h4zm-10-4h-4a2.002 2.002 0 0 1-2-2V10a2.002 2.002 0 0 1 2-2h4a2.002 2.002 0 0 1 2 2v10a2.002 2.002 0 0 1-2 2zm-4-12v10h4V10zm-4 12H5a2.002 2.002 0 0 1-2-2V10a2.002 2.002 0 0 1 2-2h6v2H5v10h6z"/>`),
		g.Group(children),
	)
}