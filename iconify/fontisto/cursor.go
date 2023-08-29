package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cursor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m0 0l13.5 12l-5.04 2.24L12 22.5L9 24l-3.593-8.4L0 18z"/>`),
		g.Group(children),
	)
}