package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Arrowup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M23.963 20.834L17.5 9.64c-.825-1.43-2.175-1.43-3 0L8.037 20.834c-.825 1.43-.15 2.598 1.5 2.598h12.926c1.65 0 2.325-1.17 1.5-2.598z"/>`),
		g.Group(children),
	)
}