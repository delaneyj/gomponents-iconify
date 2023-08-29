package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonitorO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 10.5v-8h13v8h-13ZM5 12H1a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1h-4v1.5h1.25a.75.75 0 0 1 0 1.5h-8.5a.75.75 0 0 1 0-1.5H5V12Zm1.5 0v1.5h3V12h-3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}