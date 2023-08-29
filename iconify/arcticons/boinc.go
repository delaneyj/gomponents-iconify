package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Boinc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="8.57" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5.5 5.5l7.33 7.33m19.59-2.19a15.81 15.81 0 0 0-19.59 2.19m24.53 19.59a15.81 15.81 0 0 0-2.19-19.59h0L42.5 5.5M15.58 37.36a15.81 15.81 0 0 0 19.59-2.19h0l7.33 7.33M10.64 15.58a15.81 15.81 0 0 0 2.19 19.59h0L5.5 42.5"/>`),
		g.Group(children),
	)
}