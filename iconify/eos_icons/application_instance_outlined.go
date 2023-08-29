package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApplicationInstanceOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 3v2h22V3a2.001 2.001 0 0 0-2-2H3a2.001 2.001 0 0 0-2 2Zm3 1a.987.987 0 0 1-.993-.992A1.001 1.001 0 0 1 4 2a1.013 1.013 0 0 1 1.006 1.008A.998.998 0 0 1 4 4Zm3 0a1 1 0 1 1 1-1a1.004 1.004 0 0 1-1 1ZM1 6v14a2 2 0 0 0 2 2h18a2 2 0 0 0 2-2V6Zm20 14H3V8h18Z"/><path fill="currentColor" d="M9 10v8l7-4l-7-4z"/>`),
		g.Group(children),
	)
}