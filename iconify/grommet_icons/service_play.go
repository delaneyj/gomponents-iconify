package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ServicePlay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M7 13A6 6 0 1 0 7 1a6 6 0 0 0 0 12Zm7.995 3.657a6 6 0 1 0-1.89-10.22m-8.281 6.255A6 6 0 0 0 9 23a6 6 0 0 0 2.127-11.612M6.5 6L8 7L6.5 8V6Z"/>`),
		g.Group(children),
	)
}