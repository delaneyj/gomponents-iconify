package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StopwatchSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M9 1H6V0h3v1Z"/><path fill="currentColor" fill-rule="evenodd" d="M7.5 2a6.5 6.5 0 1 0 0 13a6.5 6.5 0 0 0 0-13ZM8 8V5H7v4h3V8H8Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}