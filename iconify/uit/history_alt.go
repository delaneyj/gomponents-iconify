package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HistoryAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.728 3.18C12.31.81 6.915 2.105 4 6V3.5a.5.5 0 0 0-1 0v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 0-1H4.523a8.987 8.987 0 0 1 7.45-4A9 9 0 0 1 12 21a.5.5 0 0 0 0 1a10.005 10.005 0 0 0 8.81-5.272c2.614-4.868.786-10.934-4.082-13.547zM12 8a.5.5 0 0 0-.5.5V12a.5.5 0 0 0 .5.5h2.5a.5.5 0 0 0 0-1h-2v-3A.5.5 0 0 0 12 8z"/>`),
		g.Group(children),
	)
}