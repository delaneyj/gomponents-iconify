package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.5 20.25a.76.76 0 0 1-.75-.75v-15a.75.75 0 0 1 1.5 0v15a.76.76 0 0 1-.75.75Z"/><path fill="currentColor" d="M19.5 20.25h-15a.75.75 0 0 1 0-1.5h15a.75.75 0 0 1 0 1.5Zm-5.5-5.5a.74.74 0 0 1-.53-.22L11 12.06l-2.47 2.47a.75.75 0 0 1-1.06-1.06l3-3a.75.75 0 0 1 1.06 0L14 12.94l3.47-3.47a.75.75 0 0 1 1.06 1.06l-4 4a.74.74 0 0 1-.53.22Z"/><path fill="currentColor" d="M18.5 13.84a.76.76 0 0 1-.75-.75v-2.84H15a.75.75 0 0 1 0-1.5h3.5a.76.76 0 0 1 .75.75v3.59a.76.76 0 0 1-.75.75Z"/>`),
		g.Group(children),
	)
}