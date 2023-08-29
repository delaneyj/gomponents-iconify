package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderOuter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M56 40a16 16 0 0 0-16 16v400a16 16 0 0 0 16 16h400a16 16 0 0 0 16-16V56a16 16 0 0 0-16-16Zm384 400H72V72h368Z"/><path fill="currentColor" d="M239.272 239.272h33.455v33.455h-33.455zm0 66.91h33.455v33.454h-33.455zm0-133.819h33.455v33.454h-33.455zm0-66.909h33.455v33.454h-33.455zm0 267.637h33.455v33.454h-33.455zm66.91-133.819h33.455v33.455h-33.455zm66.909 0h33.455v33.455h-33.455zm-200.727 0h33.455v33.455h-33.455zm-66.909 0h33.455v33.455h-33.455z"/>`),
		g.Group(children),
	)
}