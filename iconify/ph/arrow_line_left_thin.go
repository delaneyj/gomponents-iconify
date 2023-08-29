package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLineLeftThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M228 128a4 4 0 0 1-4 4H81.66l65.17 65.17a4 4 0 0 1-5.66 5.66l-72-72a4 4 0 0 1 0-5.66l72-72a4 4 0 1 1 5.66 5.66L81.66 124H224a4 4 0 0 1 4 4ZM40 36a4 4 0 0 0-4 4v176a4 4 0 0 0 8 0V40a4 4 0 0 0-4-4Z"/>`),
		g.Group(children),
	)
}