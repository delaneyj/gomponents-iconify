package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashlightOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 22V10.8L1.4 4.2l1.4-1.4l18.4 18.4l-1.4 1.4l-3.8-3.8V22H8ZM6 3.15V2h12v3H7.85L6 3.15Zm10 10L9.85 7H18v1l-2 3v2.15Z"/>`),
		g.Group(children),
	)
}