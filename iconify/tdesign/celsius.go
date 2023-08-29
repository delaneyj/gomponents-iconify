package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Celsius(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.5 6a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1ZM4 6.5a2.5 2.5 0 1 1 5 0a2.5 2.5 0 0 1-5 0Zm7-.5a2 2 0 0 1 2-2h7v2h-7v12h7v2h-7a2 2 0 0 1-2-2V6Z"/>`),
		g.Group(children),
	)
}