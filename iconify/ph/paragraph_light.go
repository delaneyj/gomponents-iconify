package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParagraphLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 42H96a62 62 0 0 0 0 124h42v42a6 6 0 0 0 12 0V54h28v154a6 6 0 0 0 12 0V54h18a6 6 0 0 0 0-12Zm-70 112H96a50 50 0 0 1 0-100h42Z"/>`),
		g.Group(children),
	)
}