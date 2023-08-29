package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AzureDevops(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M15 3.622v8.512L11.5 15l-5.425-1.975v1.957l-3.071-4.01l8.951.698V4.006L15 3.622Zm-2.984.428L6.994 1v2.001L2.383 4.356L1 6.13v4.029l1.978.873V5.869l9.038-1.818Z"/>`),
		g.Group(children),
	)
}