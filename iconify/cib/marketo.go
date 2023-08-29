package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Marketo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M21.527 0v32l8.181-6.516V4.448zm-3.803 25.011l-6.584 2.875V2.782l6.584 1.948zm-15.427-.74l5.036-1.328V6.918L2.317 6.11z"/>`),
		g.Group(children),
	)
}