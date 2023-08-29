package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoWindowsFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.75 2.135v9.615h-9.5V3.492l9.5-1.357Zm-11 1.982v7.633h-8.5V5.515l8.5-1.398Zm-8.5 9.133h8.5v7.575l-8.5-.81V13.25Zm10 0h9.5v8.617l-9.5-1.385V13.25Z"/>`),
		g.Group(children),
	)
}