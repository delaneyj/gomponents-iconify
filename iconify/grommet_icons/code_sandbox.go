package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeSandbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="m12 1.5l-9 5v11l9 5l9-5v-11l-9-5Zm0 21v-11m9-5l-9 5m0 0l-9-5m18 11V12l-4.5 2.5V20l4.5-2.5Zm-18 0V12l4.5 2.5V20L3 17.5Zm9-16L7.5 4L12 6.5L16.5 4L12 1.5Z"/>`),
		g.Group(children),
	)
}