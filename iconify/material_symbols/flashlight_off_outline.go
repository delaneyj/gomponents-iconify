package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashlightOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.8 22.6L16 18.8V22H8V10.8L1.4 4.2l1.4-1.4l18.4 18.4l-1.4 1.4ZM10 20h4v-3.2l-4-4V20Zm6-6.85l-2-2v-.75l2-3V7H9.85l-2-2H16V4H6.85L6 3.15V2h12v6l-2 3v2.15Zm-4 1.65Zm0-5.65Z"/>`),
		g.Group(children),
	)
}