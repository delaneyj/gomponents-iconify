package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrightnessLow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 11a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm1.93 6.66a1 1 0 1 0 1.41 0a1 1 0 0 0-1.41 0ZM6.34 6.34a1 1 0 1 0-1.41 0a1 1 0 0 0 1.41 0ZM12 4a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm5.66 13.66a1 1 0 1 0 1.41 0a1 1 0 0 0-1.41 0ZM21 11a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm-3.34-6.07a1 1 0 1 0 1.41 0a1 1 0 0 0-1.41 0ZM12 20a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm0-14a6 6 0 1 0 6 6a6 6 0 0 0-6-6Zm0 10a4 4 0 1 1 4-4a4 4 0 0 1-4 4Z"/>`),
		g.Group(children),
	)
}