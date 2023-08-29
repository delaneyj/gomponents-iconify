package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipPrevSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.75 7a.75.75 0 0 0-1.5 0v10a.75.75 0 0 0 1.5 0V7Zm3.102 5.66a.834.834 0 0 1 0-1.32a25.009 25.009 0 0 1 6.935-3.787l.466-.165a.944.944 0 0 1 1.243.772a29.813 29.813 0 0 1 0 7.68a.944.944 0 0 1-1.243.772l-.466-.165a25.01 25.01 0 0 1-6.935-3.788Z"/>`),
		g.Group(children),
	)
}