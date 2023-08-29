package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownSquared(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v22h22V5H5zm2 2h18v18H7V7zm3.72 5.78l-1.44 1.44l6 6l.72.686l.72-.687l6-6l-1.44-1.44L16 18.064l-5.28-5.282z"/>`),
		g.Group(children),
	)
}