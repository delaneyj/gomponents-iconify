package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartPieAltSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.938 10.002a8.004 8.004 0 0 0-6.94-6.94C14.45 2.993 14 3.448 14 4v6.5a.5.5 0 0 0 .5.5H21c.552 0 1.007-.45.938-.998Z"/><path fill="currentColor" d="M12 4.5a8.5 8.5 0 1 0 8.5 8.5a.5.5 0 0 0-.5-.5h-7a.5.5 0 0 1-.5-.5V5a.5.5 0 0 0-.5-.5Z"/>`),
		g.Group(children),
	)
}