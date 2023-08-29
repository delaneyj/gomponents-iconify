package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m16 4.688l-.72.718l-11.5 11.5l1.44 1.407L16 7.53l10.78 10.782l1.44-1.406l-11.5-11.5l-.72-.718zm0 7l-.72.718l-11.5 11.5l1.44 1.407L16 14.53l10.78 10.783l1.44-1.407l-11.5-11.5l-.72-.72z"/>`),
		g.Group(children),
	)
}