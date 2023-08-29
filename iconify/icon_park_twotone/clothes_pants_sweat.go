package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClothesPantsSweat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTClothesPantsSweat0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m24 19l9 19h9L38 4H10L6 38h9l9-19Z"/><path fill="#555" d="m34 38l1 6h6v-6h-7Zm-21 6H7v-6h7l-1 6Z"/><path d="m24 4l4 7.5M24 4l-4 7.5"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTClothesPantsSweat0)"/>`),
		g.Group(children),
	)
}