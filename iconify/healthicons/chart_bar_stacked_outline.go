package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartBarStackedOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 24a1 1 0 0 0-1 1v15H8V7H6v34a1 1 0 0 0 1 1h34v-2h-2V13a1 1 0 0 0-1-1h-6a1 1 0 0 0-1 1v27h-2V21a1 1 0 0 0-1-1h-6a1 1 0 0 0-1 1v19h-2V25a1 1 0 0 0-1-1h-6Zm21 16h4V26h-4v14Zm4-16h-4V14h4v10Zm-14 6v10h4V30h-4Zm4-2h-4v-6h4v6Zm-14 6v6h4v-6h-4Zm4-2h-4v-6h4v6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}