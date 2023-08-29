package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DonutChartLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 2.05v2.012a8.001 8.001 0 1 0 5.905 14.258l1.424 1.423A9.958 9.958 0 0 1 12 22C6.477 22 2 17.523 2 12c0-5.185 3.947-9.449 9-9.95ZM21.95 13a9.954 9.954 0 0 1-2.207 5.328l-1.423-1.422A7.96 7.96 0 0 0 19.938 13h2.013ZM13.002 2.05a10.004 10.004 0 0 1 8.95 8.95h-2.013a8.004 8.004 0 0 0-6.937-6.938V2.049Z"/>`),
		g.Group(children),
	)
}