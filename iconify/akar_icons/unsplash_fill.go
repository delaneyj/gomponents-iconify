package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnsplashFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.25 2h7.5v5.625h-7.5V2ZM2 10.75h6.268v5.675h7.517V10.75H22V22H2V10.75Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}