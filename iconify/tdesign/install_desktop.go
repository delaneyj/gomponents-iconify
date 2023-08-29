package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InstallDesktop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 2v4.657l1.53-1.466l1.384 1.445L19 10.385l-3.914-3.75l1.384-1.444L18 6.657V2h2ZM1 3h13v2H3v11h18v-4h2v6H1V3Zm4 17h14v2H5v-2Z"/>`),
		g.Group(children),
	)
}