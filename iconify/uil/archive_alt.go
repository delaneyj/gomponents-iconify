package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArchiveAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 13h4a1 1 0 0 0 0-2h-4a1 1 0 0 0 0 2Zm9-10H5a3 3 0 0 0-1 5.82V18a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V8.82A3 3 0 0 0 19 3Zm-1 15a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V9h12Zm1-11H5a1 1 0 0 1 0-2h14a1 1 0 0 1 0 2Z"/>`),
		g.Group(children),
	)
}