package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieChartAltOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M6.5.5V0a.5.5 0 0 0-.5.5h.5Zm8 8V9a.5.5 0 0 0 .5-.5h-.5Zm-8 0H6V9h.5v-.5Zm0 6.5A6.5 6.5 0 0 0 13 8.5h-1A5.5 5.5 0 0 1 6.5 14v1ZM0 8.5A6.5 6.5 0 0 0 6.5 15v-1A5.5 5.5 0 0 1 1 8.5H0Zm1 0A5.5 5.5 0 0 1 6.5 3V2A6.5 6.5 0 0 0 0 8.5h1ZM6.5 1A7.5 7.5 0 0 1 14 8.5h1A8.5 8.5 0 0 0 6.5 0v1ZM6 .5v8h1v-8H6ZM6.5 9h8V8h-8v1Z"/>`),
		g.Group(children),
	)
}