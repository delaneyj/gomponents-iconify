package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Info(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 10a6 6 0 1 0 12 0a6 6 0 0 0-12 0m6-8a8 8 0 1 1 0 16a8 8 0 0 1 0-16m1 7v5H9V9zm0-1V6H9v2z"/>`),
		g.Group(children),
	)
}