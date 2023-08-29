package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryShareSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 17v-5h6.175L14.6 10.4L16 9l4 4l-4 4l-1.425-1.425l1.6-1.575H12v3h-2Zm-3 5V4h3V2h4v2h3v4h-2V6H9v14h6v-2h2v4H7Z"/>`),
		g.Group(children),
	)
}