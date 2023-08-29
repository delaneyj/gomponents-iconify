package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderFull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8 10h10v2H8zm0 5h6v2H8z"/><path fill="currentColor" d="M29 29H3V3h26ZM5 27h22V5H5Z"/>`),
		g.Group(children),
	)
}