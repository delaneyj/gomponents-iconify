package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleLeftFortyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.5 24a5 5 0 1 1 10 0a5 5 0 0 1-10 0ZM4 24c0-5.523 4.477-10 10-10h20c5.523 0 10 4.477 10 10s-4.477 10-10 10H14C8.477 34 4 29.523 4 24Zm10-7.5a7.5 7.5 0 0 0 0 15h20a7.5 7.5 0 0 0 0-15H14Z"/>`),
		g.Group(children),
	)
}