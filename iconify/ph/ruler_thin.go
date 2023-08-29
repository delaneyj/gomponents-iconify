package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RulerThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232.49 76.2L179.8 23.51a12 12 0 0 0-17 0L23.51 162.83a12 12 0 0 0 0 17l52.69 52.66a12 12 0 0 0 17 0L232.49 93.17a12 12 0 0 0 0-16.97Zm-5.66 11.31L87.51 226.83a4 4 0 0 1-5.65 0l-52.69-52.69a4 4 0 0 1 0-5.65L64 133.66l29.17 29.17a4 4 0 1 0 5.66-5.66L69.65 128L96 101.66l29.17 29.17a4 4 0 0 0 5.66-5.66L101.65 96L128 69.66l29.17 29.17a4 4 0 1 0 5.66-5.66L133.66 64l34.83-34.83a4 4 0 0 1 5.65 0l52.69 52.69a4 4 0 0 1 0 5.65Z"/>`),
		g.Group(children),
	)
}