package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M4.5 6.5L1.328 9.672a2.828 2.828 0 1 0 4 4L8.5 10.5m2-2l3.172-3.172a2.829 2.829 0 0 0-4-4L6.5 4.5m-2 6l6-6"/>`),
		g.Group(children),
	)
}