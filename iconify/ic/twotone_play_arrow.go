package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotonePlayArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 8.64v6.72L15.27 12z" opacity=".3"/><path fill="currentColor" d="m8 19l11-7L8 5v14zm2-10.36L15.27 12L10 15.36V8.64z"/>`),
		g.Group(children),
	)
}