package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortDescending(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M40 128a8 8 0 0 1 8-8h72a8 8 0 0 1 0 16H48a8 8 0 0 1-8-8Zm8-56h56a8 8 0 0 0 0-16H48a8 8 0 0 0 0 16Zm136 112H48a8 8 0 0 0 0 16h136a8 8 0 0 0 0-16Zm45.66-101.66l-40-40a8 8 0 0 0-11.32 0l-40 40a8 8 0 0 0 11.32 11.32L176 67.31V144a8 8 0 0 0 16 0V67.31l26.34 26.35a8 8 0 0 0 11.32-11.32Z"/>`),
		g.Group(children),
	)
}