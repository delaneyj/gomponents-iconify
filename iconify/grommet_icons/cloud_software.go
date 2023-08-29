package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudSoftware(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M8 23h8V12H8v11Zm0-7h8m-4-4v4M6 6V5c0-3 1.5-4 4-4h4c2.5 0 4 1.5 4 4v1c3 0 5 2 5 5s-2 5-5 5M14 6H6c-3 0-5 1.5-5 5s2 5 5 5"/>`),
		g.Group(children),
	)
}