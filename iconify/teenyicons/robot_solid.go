package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RobotSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M5 8.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0Zm4 0a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0Z"/><path fill="currentColor" fill-rule="evenodd" d="M8 2.022A5.5 5.5 0 0 1 13 7.5v6a1.5 1.5 0 0 1-1.5 1.5h-8A1.5 1.5 0 0 1 2 13.5v-6a5.5 5.5 0 0 1 5-5.478V0h1v2.022ZM5.5 7a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Zm4 0a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Zm1.5 5H4v-1h7v1Z" clip-rule="evenodd"/><path fill="currentColor" d="M0 8v4h1V8H0Zm15 0h-1v4h1V8Z"/>`),
		g.Group(children),
	)
}