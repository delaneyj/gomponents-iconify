package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RouteOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 19a2 2 0 1 0 4 0a2 2 0 1 0-4 0M16 5a2 2 0 1 0 4 0a2 2 0 1 0-4 0m-4 14h4.5c.71 0 1.372-.212 1.924-.576m1.545-2.459A3.5 3.5 0 0 0 16.5 12h-.499m-4 0H8.5a3.5 3.5 0 0 1-2.477-5.972M8.5 5H12M3 3l18 18"/>`),
		g.Group(children),
	)
}