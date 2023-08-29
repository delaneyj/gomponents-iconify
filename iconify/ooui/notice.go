package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Notice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 0a10 10 0 1 0 10 10A10 10 0 0 0 10 0zm1 16H9v-2h2zm0-4H9V4h2z"/>`),
		g.Group(children),
	)
}