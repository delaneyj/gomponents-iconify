package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoltOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.55 18.2l5.175-6.2h-4l.725-5.675L7.825 13H11.3l-.75 5.2ZM8 22l1-7H4l9-13h2l-1 8h6L10 22H8Zm3.775-9.75Z"/>`),
		g.Group(children),
	)
}