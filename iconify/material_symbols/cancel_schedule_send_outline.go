package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CancelScheduleSendOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.225 19.475l1.75-1.75l1.75 1.75l.75-.75L17.75 17l1.75-1.75l-.75-.75L17 16.25l-1.75-1.75l-.75.75L16.25 17l-1.75 1.75l.725.725ZM3 20V4l14.3 6H17q-.875 0-1.65.2t-1.5.55L5 7v3.5l6 1.5l-6 1.5V17l5.4-2.3q-.2.575-.3 1.137T10 17v.05L3 20Zm14 2q-2.075 0-3.538-1.463T12 17q0-2.075 1.463-3.538T17 12q2.075 0 3.538 1.463T22 17q0 2.075-1.463 3.538T17 22ZM5 14.7V7v10v-2.3Z"/>`),
		g.Group(children),
	)
}