package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditFlipLeftFlipLeftObjectWork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 13.5V.5a6.5 6.5 0 0 0 0 13ZM9.61 1.05a6.13 6.13 0 0 1 1.3.76a6.34 6.34 0 0 1 1.09 1M9.61 13a6.13 6.13 0 0 0 1.3-.76a6.34 6.34 0 0 0 1.09-1m1.33-5.74a6.7 6.7 0 0 1 0 3"/>`),
		g.Group(children),
	)
}