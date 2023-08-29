package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gift(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M3 11h18v12H3V11Zm-1 0V7h20v4H2Zm10 12V7v16ZM7 7h5s-2-5-5-5C3.5 2 3 7 7 7Zm10.184 0h-5s1.816-5 5-5c3.316 0 4 5 0 5Z"/>`),
		g.Group(children),
	)
}