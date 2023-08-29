package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkMultipleTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1 8a4 4 0 0 1 4-4h6a4 4 0 0 1 0 8H9.5a.5.5 0 0 1 0-1H11a3 3 0 1 0 0-6H5a3 3 0 0 0-.872 5.871a5.007 5.007 0 0 0-.126 1.003A4.002 4.002 0 0 1 1 8Zm17 4a3.001 3.001 0 0 0-2.128-2.871c.075-.324.118-.66.127-1.003A4.002 4.002 0 0 1 15 16H9a4 4 0 0 1 0-8h1.5a.5.5 0 0 1 0 1H9a3 3 0 1 0 0 6h6a3 3 0 0 0 3-3Z"/>`),
		g.Group(children),
	)
}