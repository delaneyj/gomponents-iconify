package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cointrend(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.91 28.92v.118c0 5.23-4.436 9.472-9.91 9.472h0c-5.474 0-9.91-4.241-9.91-9.472v-9.65c0-5.232 4.436-9.473 9.91-9.473h0c5.474 0 9.91 4.24 9.91 9.472v.117M24.085 44.5v-5.66m-.122-29.68V3.5"/>`),
		g.Group(children),
	)
}