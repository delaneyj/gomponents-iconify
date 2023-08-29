package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldDismissTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.277 2.084a15.05 15.05 0 0 0 6.294 2.421A.5.5 0 0 1 17 5v4.5c0 3.891-2.307 6.73-6.82 8.467a.5.5 0 0 1-.36 0C5.308 16.23 3 13.39 3 9.5V5a.5.5 0 0 1 .43-.495a15.05 15.05 0 0 0 6.293-2.421a.5.5 0 0 1 .554 0ZM8.03 6.97a.75.75 0 0 0-1.06 1.06L8.94 10l-1.97 1.97a.75.75 0 1 0 1.06 1.06L10 11.06l1.97 1.97a.75.75 0 1 0 1.06-1.06L11.06 10l1.97-1.97a.75.75 0 0 0-1.06-1.06L10 8.94L8.03 6.97Z"/>`),
		g.Group(children),
	)
}