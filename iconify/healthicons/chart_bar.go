package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartBar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" d="M12 24a1 1 0 0 0-1 1v15H8V7H6v34a1 1 0 0 0 1 1h34v-2h-2V13a1 1 0 0 0-1-1h-6a1 1 0 0 0-1 1v27h-2V21a1 1 0 0 0-1-1h-6a1 1 0 0 0-1 1v19h-2V25a1 1 0 0 0-1-1h-6Z"/>`),
		g.Group(children),
	)
}