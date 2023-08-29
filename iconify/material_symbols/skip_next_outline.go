package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipNextOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.5 18V6h2v12h-2Zm-11 0V6l9 6l-9 6Zm2-6Zm0 2.25L10.9 12L7.5 9.75v4.5Z"/>`),
		g.Group(children),
	)
}