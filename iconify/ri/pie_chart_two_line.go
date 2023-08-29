package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieChartTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 .5C18.351.5 23.5 5.649 23.5 12c0 .337-.015.67-.043 1h-1.506c-.502 5.053-4.766 9-9.951 9c-5.523 0-10-4.477-10-10c0-5.185 3.947-9.449 9-9.95V.542c.33-.029.663-.043 1-.043Zm-1 3.562A8.001 8.001 0 0 0 12 20a8.001 8.001 0 0 0 7.938-7H11V4.062Zm2-1.51V11h8.448A9.503 9.503 0 0 0 13 2.552Z"/>`),
		g.Group(children),
	)
}