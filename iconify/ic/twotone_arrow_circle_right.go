package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneArrowCircleRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 12c0-4.41 3.59-8 8-8s8 3.59 8 8s-3.59 8-8 8s-8-3.59-8-8m8-1H8v2h4v3l4-4l-4-4v3z" opacity=".3"/><path fill="currentColor" d="M4 12c0-4.41 3.59-8 8-8s8 3.59 8 8s-3.59 8-8 8s-8-3.59-8-8m-2 0c0 5.52 4.48 10 10 10s10-4.48 10-10S17.52 2 12 2S2 6.48 2 12zm10-1H8v2h4v3l4-4l-4-4v3z"/>`),
		g.Group(children),
	)
}