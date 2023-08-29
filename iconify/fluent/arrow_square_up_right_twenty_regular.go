package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSquareUpRightTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 3a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3H6ZM4 6a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6Zm8.691 1.038c.058.024.113.06.16.106l.005.005A.5.5 0 0 1 13 7.5v5a.5.5 0 0 1-1 0V8.707l-4.146 4.147a.5.5 0 0 1-.708-.708L11.293 8H7.5a.5.5 0 0 1 0-1h5c.068 0 .132.013.191.038Z"/>`),
		g.Group(children),
	)
}