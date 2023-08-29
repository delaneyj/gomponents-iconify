package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DesktopMacSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 4a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2h-2v.5a.5.5 0 0 0 .5.5h.5a.5.5 0 0 1 0 1H5a.5.5 0 0 1 0-1h.5a.5.5 0 0 0 .5-.5V12H4a2 2 0 0 1-2-2V4Zm1 6a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1H3Zm10-1V4a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v5h10Zm-4 3.5V12H7v.5c0 .175-.03.344-.085.5H9.09a1.5 1.5 0 0 1-.09-.5Z"/>`),
		g.Group(children),
	)
}