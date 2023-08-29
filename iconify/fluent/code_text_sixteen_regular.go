package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeTextSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1.5 3a.5.5 0 0 0 0 1H7a.5.5 0 0 0 0-1H1.5Zm3 3a.5.5 0 0 0 0 1h4a.5.5 0 0 0 0-1h-4ZM3 9.5a.5.5 0 0 1 .5-.5h8a.5.5 0 0 1 0 1h-8a.5.5 0 0 1-.5-.5ZM1.5 12a.5.5 0 0 0 0 1H9a.5.5 0 0 0 0-1H1.5Zm9-5.5A.5.5 0 0 1 11 6h2.5a.5.5 0 0 1 0 1H11a.5.5 0 0 1-.5-.5ZM9.5 3a.5.5 0 0 0 0 1h5a.5.5 0 0 0 0-1h-5Z"/>`),
		g.Group(children),
	)
}