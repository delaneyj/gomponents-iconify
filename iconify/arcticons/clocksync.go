package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clocksync(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="17" cy="24" r="12.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.442 17.127a9.5 9.5 0 1 1 0 13.746"/><circle cx="17" cy="24" r="2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="34" cy="24" r="2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17 22v-8.196M19 24h5.501m10.905-1.406l3.8-3.8m-6.618 3.794l-2.065-2.065"/>`),
		g.Group(children),
	)
}