package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16.867 10.837a.5.5 0 1 0-.739-.674l-5.63 6.168V2.5a.5.5 0 0 0-1 0v13.828l-5.629-6.165a.5.5 0 0 0-.738.674l6.314 6.916a.747.747 0 0 0 1.108 0l6.314-6.916Z"/>`),
		g.Group(children),
	)
}