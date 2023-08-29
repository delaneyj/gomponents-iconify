package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FireSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.063 2.065a.667.667 0 0 0-.759.149a.805.805 0 0 0-.205.738c.099.44.151.9.151 1.37c0 1.076-.271 1.83-.738 2.455c-.474.635-1.16 1.152-2.027 1.735a.25.25 0 0 0-.038.031l-.105.106a6.75 6.75 0 1 0 9.685 2.63a.25.25 0 0 0-.413-.05l-.208.241c-.878 1.026-1.587 1.855-3.04 2.225c-.062.015-.1.004-.127-.013a.244.244 0 0 1-.091-.124a.411.411 0 0 1 .074-.414c.777-.843 1.329-1.987 1.526-3.614c.37-3.048-1.015-6.294-3.685-7.465Z"/>`),
		g.Group(children),
	)
}