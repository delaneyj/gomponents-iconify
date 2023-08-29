package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThermometerProbeOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m2 3.64l7.24 7.24l-6.43 6.43a2.758 2.758 0 0 0 0 3.89a2.758 2.758 0 0 0 3.89 0l6.43-6.43l7.6 7.6L22 21.1L3.27 2.37L2 3.64m8.23 8.22l1.91 1.91a1.376 1.376 0 1 1-1.91-1.91m4.23-.86L13 9.55l.47-.08l6-6L22 2l-1.46 2.54l-6 6l-.08.46Z"/>`),
		g.Group(children),
	)
}