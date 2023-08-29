package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryShare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 22q-.425 0-.713-.288T7 21V5q0-.425.288-.713T8 4h2V2h4v2h2q.425 0 .713.288T17 5v3h-2V6H9v14h6v-2h2v3q0 .425-.288.713T16 22H8Zm2-5v-3q0-.825.588-1.413T12 12h4.175L14.6 10.4L16 9l4 4l-4 4l-1.425-1.425l1.6-1.575H12v3h-2Z"/>`),
		g.Group(children),
	)
}