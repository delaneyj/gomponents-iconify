package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentTextSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 4.5V1H4.5A1.5 1.5 0 0 0 3 2.5v11A1.5 1.5 0 0 0 4.5 15h7a1.5 1.5 0 0 0 1.5-1.5V6H9.5A1.5 1.5 0 0 1 8 4.5Zm1 0V1.25L12.75 5H9.5a.5.5 0 0 1-.5-.5ZM5.5 8h5a.5.5 0 0 1 0 1h-5a.5.5 0 0 1 0-1ZM5 10.5a.5.5 0 0 1 .5-.5h5a.5.5 0 0 1 0 1h-5a.5.5 0 0 1-.5-.5Zm.5 1.5h5a.5.5 0 0 1 0 1h-5a.5.5 0 0 1 0-1Z"/>`),
		g.Group(children),
	)
}