package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ctrl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 0v16h16V0H0zm4.19 12C2.16 12 1 10.54 1 8s1.16-4 3.19-4h.029c.539 0 1.052.114 1.515.32l-.424.901a2.719 2.719 0 0 0-1.08-.22h-.042C2.38 5.001 2 6.631 2 8.001s.38 3 2.19 3c.497-.013.96-.145 1.366-.368l.444.898a3.943 3.943 0 0 1-1.806.47zM9 7H8v3.5a.902.902 0 0 0 1.005.499L8.999 12a1.822 1.822 0 0 1-1.998-1.428L6.999 7h-.51V6h.51V5h1v1h1v1zm2 2v3h-1V6h1v.92a2.447 2.447 0 0 1 2.005-.919l-.004 1a1.88 1.88 0 0 0-2 2.006zm4 3h-1V3h1v9z"/>`),
		g.Group(children),
	)
}