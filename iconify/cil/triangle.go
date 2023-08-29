package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Triangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M273.342 66.038a20 20 0 0 0-34.684 0L29.569 430.007a20 20 0 0 0 17.342 29.963h418.178a20 20 0 0 0 17.342-29.962ZM67.644 427.97L256 100.091L444.356 427.97Z"/>`),
		g.Group(children),
	)
}