package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowThickToLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M336 176V56h-38.627l-200 199.8L297.363 456H336V336h160V176Zm128 128H304v113.361L142.625 255.826L304 94.616V208h160ZM16 56h32v400H16z"/>`),
		g.Group(children),
	)
}