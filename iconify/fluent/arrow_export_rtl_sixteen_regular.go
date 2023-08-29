package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowExportRtlSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14.5 3a.5.5 0 0 0-.5.5V12a.5.5 0 0 0 1 0V3.5a.5.5 0 0 0-.5-.5Zm-9.146.646a.5.5 0 0 0-.708 0l-3.5 3.5a.5.5 0 0 0 0 .708l3.5 3.5a.5.5 0 0 0 .708-.708L2.707 8H11.5a.5.5 0 0 0 0-1H2.707l2.647-2.646a.5.5 0 0 0 0-.708Z"/>`),
		g.Group(children),
	)
}