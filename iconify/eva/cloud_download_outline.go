package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudDownloadOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaCloudDownloadOutline0"><g id="evaCloudDownloadOutline1"><g id="evaCloudDownloadOutline2" fill="currentColor"><path d="M14.31 16.38L13 17.64V12a1 1 0 0 0-2 0v5.59l-1.29-1.3a1 1 0 0 0-1.42 1.42l3 3A1 1 0 0 0 12 21a1 1 0 0 0 .69-.28l3-2.9a1 1 0 1 0-1.38-1.44Z"/><path d="M17.67 7A6 6 0 0 0 6.33 7a5 5 0 0 0-3.08 8.27A1 1 0 1 0 4.75 14A3 3 0 0 1 7 9h.1a1 1 0 0 0 1-.8a4 4 0 0 1 7.84 0a1 1 0 0 0 1 .8H17a3 3 0 0 1 2.25 5a1 1 0 0 0 .09 1.42a1 1 0 0 0 .66.25a1 1 0 0 0 .75-.34A5 5 0 0 0 17.67 7Z"/></g></g></g>`),
		g.Group(children),
	)
}