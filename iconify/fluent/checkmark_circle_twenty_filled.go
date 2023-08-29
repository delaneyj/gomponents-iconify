package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckmarkCircleTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 2a8 8 0 1 1 0 16a8 8 0 0 1 0-16Zm3.358 5.646a.5.5 0 0 0-.637-.057l-.07.057L9 11.298L7.354 9.651l-.07-.058a.5.5 0 0 0-.695.696l.057.07l2 2l.07.057a.5.5 0 0 0 .568 0l.07-.058l4.004-4.004l.058-.07a.5.5 0 0 0-.058-.638Z"/>`),
		g.Group(children),
	)
}