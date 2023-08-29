package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartBarStacked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7 42a1 1 0 0 1-1-1V7h2v33h3V25a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v15h2V21a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v19h2V13a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v27h2v2H7Zm30-18V14h-4v10h4Zm-14 4h4v-6h-4v6Zm-6 4v-6h-4v6h4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}