package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRightTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.837 3.13a.5.5 0 0 0-.674.74L16.33 9.5H2.5a.5.5 0 0 0 0 1h13.828l-6.165 5.628a.5.5 0 0 0 .674.739l6.916-6.314a.747.747 0 0 0 0-1.108L10.837 3.13Z"/>`),
		g.Group(children),
	)
}