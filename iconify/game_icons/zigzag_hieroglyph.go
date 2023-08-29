package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZigzagHieroglyph(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M204.3 118.4L152.5 188l-51.3-69.1L20.99 220l21.94 17.4l57.47-72.5l52.1 69.9l51.8-69.4l51.6 69.5l51.7-69.6l51.8 69.6l52-70l57.6 72.5l22-17.4l-80.4-101.1l-51.2 69l-51.8-69.4l-51.7 69.4l-51.6-69.5zm0 156.4l-51.8 69.3l-51.3-68.8l-80.17 100.8l21.92 17.4l57.55-72.4l52 69.8l51.8-69.3l51.6 69.3l51.7-69.3l51.8 69.3l52-69.8l57.7 72.4l21.8-17.4l-80.3-100.8l-51.2 68.8l-51.8-69.3l-51.7 69.3l-51.6-69.3z"/>`),
		g.Group(children),
	)
}