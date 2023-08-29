package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Analogue(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M7.292 17.073c-9.131-.573-9.989 13.12-.86 13.692c9.131.573 9.989-13.12.86-13.692zm23.787 3.416L21.932 4.645c-4.577-7.921-16.459-1.057-11.88 6.865l9.145 15.839a6.871 6.871 0 0 0 9.376 2.515a6.858 6.858 0 0 0 2.505-9.375z"/>`),
		g.Group(children),
	)
}