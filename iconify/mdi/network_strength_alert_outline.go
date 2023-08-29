package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkStrengthAlertOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M21.001 1l-20 20.001h16.002V19H6L19 5.83v1.173H21M19 8.998v8.004H21V8.999m-2.002 10V21H21V19" fill="currentColor"/>`),
		g.Group(children),
	)
}