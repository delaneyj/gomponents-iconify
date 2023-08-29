package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kenya(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M410.7 119.7v182.1l27.4 25.9l-105.9 154.6l-73.9-41.8l-5.4-39.7L74.66 296.4l36.64-29L62.47 253l50.33-78.4l-.8-61.8l-47.25-60.96l38.15-19.05l99-3.05S307.1 83.8 310.1 83.09c3.1-.81 91.5-36.58 91.5-36.58l47.9 23.61z"/>`),
		g.Group(children),
	)
}