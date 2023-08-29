package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataPieTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 12a1 1 0 0 1-1-1V5.083A6.002 6.002 0 0 0 9 17a6.002 6.002 0 0 0 5.917-5H9Zm-.997-7.93C8.55 3.992 9 4.448 9 5v6h6c.552 0 1.008.45.93.997A7.002 7.002 0 0 1 2 11a7.002 7.002 0 0 1 6.003-6.93ZM11 9V3a6 6 0 0 1 6 6h-6Zm6.062 1c.498 0 .927-.366.937-.864L18 9a7 7 0 0 0-7.136-6.999c-.498.01-.864.44-.864.937V9a1 1 0 0 0 1 1h6.062Z"/>`),
		g.Group(children),
	)
}