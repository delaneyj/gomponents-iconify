package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeviceCare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.204 43.628A20 20 0 0 1 28.562 9.974"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.497 10.122a20 20 0 0 1 11.18 5.428M31.4 22.336a8.5 8.5 0 1 1-8.167 8.167m.16-3.818l-.16 3.818"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m18.457 30.674l4.936-3.989l4.935 4.057"/>`),
		g.Group(children),
	)
}