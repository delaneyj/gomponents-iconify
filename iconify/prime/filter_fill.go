package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.17 3.91a.76.76 0 0 0-.67-.41h-15a.76.76 0 0 0-.67.41a.73.73 0 0 0 .07.78L9.25 12v7.75a.76.76 0 0 0 .75.75h4a.76.76 0 0 0 .75-.75V12l5.35-7.31a.73.73 0 0 0 .07-.78Z"/>`),
		g.Group(children),
	)
}