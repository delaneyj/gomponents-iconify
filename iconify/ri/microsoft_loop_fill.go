package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrosoftLoopFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 12c0 5.523-4.477 10-10 10H6.665A7.99 7.99 0 0 0 10 15.5h2A3.5 3.5 0 1 0 8.5 12v3.5A6.5 6.5 0 0 1 2 22V12C2 6.477 6.477 2 12 2s10 4.477 10 10Z"/>`),
		g.Group(children),
	)
}