package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QrScanFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 15v5.007a.994.994 0 0 1-.993.993H3.993A.993.993 0 0 1 3 20.007V15h18ZM2 11h20v2H2v-2Zm19-2H3V3.993C3 3.445 3.445 3 3.993 3h16.014c.548 0 .993.445.993.993V9Z"/>`),
		g.Group(children),
	)
}