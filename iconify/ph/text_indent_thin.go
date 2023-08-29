package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextIndentThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M220 128a4 4 0 0 1-4 4H112a4 4 0 0 1 0-8h104a4 4 0 0 1 4 4ZM112 68h104a4 4 0 0 0 0-8H112a4 4 0 0 0 0 8Zm104 120H40a4 4 0 0 0 0 8h176a4 4 0 0 0 0-8ZM37.17 138.83a4 4 0 0 0 5.66 0l40-40a4 4 0 0 0 0-5.66l-40-40a4 4 0 0 0-5.66 5.66L74.34 96l-37.17 37.17a4 4 0 0 0 0 5.66Z"/>`),
		g.Group(children),
	)
}