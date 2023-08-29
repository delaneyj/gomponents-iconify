package nonicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScalaSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M12.712.261a.765.765 0 0 1 .288.6v12.123c0 .35-.234.653-.568.738l-8.5 2.155a.742.742 0 0 1-.644-.138a.765.765 0 0 1-.288-.6V3.016c0-.35.234-.653.568-.738l8.5-2.155a.742.742 0 0 1 .644.138ZM4.5 3.61v2.323l7-1.886V1.835l-7 1.775Zm7 2.012l-7 1.885v2.617l7-1.775V5.622Zm0 4.295l-7 1.775v2.473l7-1.775V9.917Z"/>`),
		g.Group(children),
	)
}