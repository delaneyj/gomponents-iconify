package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ErrorCircleTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 2a8 8 0 1 1 0 16a8 8 0 0 1 0-16Zm0 10.5a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5ZM10 6a.5.5 0 0 0-.492.41L9.5 6.5V11l.008.09a.5.5 0 0 0 .984 0L10.5 11V6.5l-.008-.09A.5.5 0 0 0 10 6Z"/>`),
		g.Group(children),
	)
}