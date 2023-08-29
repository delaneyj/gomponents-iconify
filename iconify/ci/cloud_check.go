package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 11l-4 4l-2-2m14 2a4 4 0 0 0-4.07-4A7.002 7.002 0 0 0 5.669 9.01A5 5 0 0 0 6 19h13a4 4 0 0 0 4-4Z"/>`),
		g.Group(children),
	)
}