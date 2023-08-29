package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserFocusFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 48v28a8 8 0 0 1-16 0V48h-28a8 8 0 0 1 0-16h28a16 16 0 0 1 16 16Zm-8 124a8 8 0 0 0-8 8v28h-28a8 8 0 0 0 0 16h28a16 16 0 0 0 16-16v-28a8 8 0 0 0-8-8ZM76 208H48v-28a8 8 0 0 0-16 0v28a16 16 0 0 0 16 16h28a8 8 0 0 0 0-16ZM40 84a8 8 0 0 0 8-8V48h28a8 8 0 0 0 0-16H48a16 16 0 0 0-16 16v28a8 8 0 0 0 8 8Zm61 57.51a67.94 67.94 0 0 0-27.4 21.68A8 8 0 0 0 80 176h96a8 8 0 0 0 6.4-12.81a67.94 67.94 0 0 0-27.4-21.68a40 40 0 1 0-53.94 0Z"/>`),
		g.Group(children),
	)
}