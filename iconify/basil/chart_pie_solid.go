package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartPieSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.25 2.745c0-.116-.1-.208-.216-.196A9.5 9.5 0 0 0 2.5 12a9.5 9.5 0 0 0 16.243 6.692a.197.197 0 0 0-.017-.292l-7.197-5.817A.75.75 0 0 1 11.25 12V2.745Z"/><path fill="currentColor" d="M19.67 17.234c.09.073.224.054.288-.044a9.446 9.446 0 0 0 1.494-4.225a.197.197 0 0 0-.197-.215h-6.568a.2.2 0 0 0-.126.355l5.108 4.13Zm1.585-5.984c.116 0 .208-.1.197-.216a9.503 9.503 0 0 0-8.486-8.486a.197.197 0 0 0-.216.197v8.205a.3.3 0 0 0 .3.3h8.205Z"/>`),
		g.Group(children),
	)
}