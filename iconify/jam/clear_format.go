package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClearFormat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.326 2L3.95 9.309a1 1 0 1 1-1.902-.618L4.223 2H1a1 1 0 1 1 0-2h8a1 1 0 1 1 0 2H6.326zm3.33 6.243l.708.707a1 1 0 1 1-1.414 1.414l-.707-.707l-.707.707A1 1 0 0 1 6.12 8.95l.707-.707l-.707-.707A1 1 0 0 1 7.536 6.12l.707.707l.707-.707a1 1 0 1 1 1.414 1.415l-.707.707z"/>`),
		g.Group(children),
	)
}