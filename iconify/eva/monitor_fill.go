package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonitorFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaMonitorFill0"><g id="evaMonitorFill1"><path id="evaMonitorFill2" fill="currentColor" d="M19 3H5a3 3 0 0 0-3 3v5h20V6a3 3 0 0 0-3-3ZM2 14a3 3 0 0 0 3 3h6v2H7a1 1 0 0 0 0 2h10a1 1 0 0 0 0-2h-4v-2h6a3 3 0 0 0 3-3v-1H2Z"/></g></g>`),
		g.Group(children),
	)
}