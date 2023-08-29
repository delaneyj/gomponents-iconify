package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paragraph(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M440 64H208a128 128 0 0 0 0 256h56v128h112V96h64ZM264 288h-56a96 96 0 0 1 0-192h56Zm80 128h-48V96h48Z"/>`),
		g.Group(children),
	)
}