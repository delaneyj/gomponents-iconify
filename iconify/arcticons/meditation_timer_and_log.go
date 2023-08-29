package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MeditationTimerAndLog(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="29.585" cy="25.137" r="5.635" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.22 25.137c0-8.501-5.239-11.17-9.885-11.17s-11.368 2.372-11.368 11.17s4.448 14.827 15.619 14.827c10.577 0 15.915-8.303 15.915-14.827C45.5 18.909 44.577 2.5 24.18 2.5C6.553 2.5 2.5 17.92 2.5 25.137a30.077 30.077 0 0 0 7.513 18.386"/>`),
		g.Group(children),
	)
}