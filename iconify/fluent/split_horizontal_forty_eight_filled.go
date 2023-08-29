package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitHorizontalFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M42.75 25.5a1.25 1.25 0 1 0 0-2.5H5.25a1.25 1.25 0 1 0 0 2.5h37.5ZM8 39.75V27.5h32v12.25A4.25 4.25 0 0 1 35.75 44h-23.5A4.25 4.25 0 0 1 8 39.75ZM40 21V8.25A4.25 4.25 0 0 0 35.75 4h-23.5A4.25 4.25 0 0 0 8 8.25V21h32Z"/>`),
		g.Group(children),
	)
}