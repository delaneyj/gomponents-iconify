package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CancelScheduleSend(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.225 19.475l1.75-1.75l1.75 1.75l.75-.75L17.75 17l1.75-1.75l-.75-.75L17 16.25l-1.75-1.75l-.75.75L16.25 17l-1.75 1.75l.725.725ZM3 20v-6l8-2l-8-2V4l14.3 6H17q-2.925 0-4.963 2.063T10 17.05L3 20Zm14 2q-2.075 0-3.538-1.463T12 17q0-2.075 1.463-3.538T17 12q2.075 0 3.538 1.463T22 17q0 2.075-1.463 3.538T17 22Z"/>`),
		g.Group(children),
	)
}