package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Industry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M14 1v12H1V8.72a.5.5 0 0 1 .17-.37l3-3.22a.5.5 0 0 1 .83.38v3l3.16-3.37a.5.5 0 0 1 .84.37V11h3V1h2z"/>`),
		g.Group(children),
	)
}