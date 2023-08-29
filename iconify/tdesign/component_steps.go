package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComponentSteps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 12a3.5 3.5 0 0 1 6.855-1h.79a3.502 3.502 0 0 1 6.71 0h.79A3.502 3.502 0 0 1 23 12a3.5 3.5 0 0 1-6.855 1h-.79a3.502 3.502 0 0 1-6.71 0h-.79A3.502 3.502 0 0 1 1 12Zm3.5-1.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Zm7.5 0a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Zm7.5 0a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Z"/>`),
		g.Group(children),
	)
}