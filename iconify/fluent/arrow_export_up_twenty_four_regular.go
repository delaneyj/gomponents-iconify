package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowExportUpTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.28 2.22a.75.75 0 0 0-1.06 0l-5 5a.75.75 0 0 0 1.06 1.06L11 4.56v13.69a.75.75 0 0 0 1.5 0V4.56l3.72 3.72a.75.75 0 1 0 1.06-1.06l-5-5ZM5.25 20.5a.75.75 0 0 0 0 1.5h13a.75.75 0 0 0 0-1.5h-13Z"/>`),
		g.Group(children),
	)
}