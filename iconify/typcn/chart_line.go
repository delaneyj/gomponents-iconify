package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.75 15.561a2 2 0 0 0 2.811-.313l2.789-3.486L12.8 13.6a2 2 0 0 0 2.762-.352l4-5a2 2 0 0 0-3.124-2.499l-2.789 3.486L11.2 7.4a2 2 0 0 0-2.762.35l-4 5a2.001 2.001 0 0 0 .312 2.811zM5 21h14a1 1 0 1 0 0-2H5a1 1 0 1 0 0 2z"/>`),
		g.Group(children),
	)
}