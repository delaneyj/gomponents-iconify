package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gitter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8.15 20.156H5V0h3.15zm6.281-15.4h-3.15V32h3.15zm6.288 0h-3.15V32h3.15zM27 4.75h-3.15v15.438H27z"/>`),
		g.Group(children),
	)
}