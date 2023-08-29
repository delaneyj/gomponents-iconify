package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartLineStackedOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" d="M6 41a1 1 0 0 0 1 1h34v-2H8v-.638l9.331-11.198l11.353 3.785a1 1 0 0 0 1.084-.309l10-12l-1.536-1.28l-9.563 11.476l-11.353-3.785a1 1 0 0 0-1.084.309L8 36.238v-6.851l9.289-10.218l10.37 3.77a1 1 0 0 0 1.11-.299l10-12l-1.537-1.28l-9.55 11.46l-10.34-3.76a1 1 0 0 0-1.082.267L8 26.413V7H6v34Z"/>`),
		g.Group(children),
	)
}