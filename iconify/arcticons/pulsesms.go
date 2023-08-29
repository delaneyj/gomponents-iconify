package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pulsesms(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.335 14.005h23.128m-23.128 6.291h16.027m-16.027 6.292h8.926M40.858 5.5H7.142a2.006 2.006 0 0 0-2 2v35l7.418-7.419h28.298a2.006 2.006 0 0 0 2-2V7.5a2.006 2.006 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}