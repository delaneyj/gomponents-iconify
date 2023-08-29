package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NewSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.502 11h5a.5.5 0 0 1 .09.992l-.09.008h-5a.5.5 0 0 1-.09-.992l.09-.008h5h-5Zm1.646-5.854a.5.5 0 0 1 .638-.057l.07.058l3.998 4a.5.5 0 0 1-.638.764l-.07-.058l-3.998-4a.5.5 0 0 1 0-.707ZM11.502 3a.5.5 0 0 1 .491.41l.009.09v5a.5.5 0 0 1-.992.09l-.008-.09v-5a.5.5 0 0 1 .5-.5Z"/>`),
		g.Group(children),
	)
}