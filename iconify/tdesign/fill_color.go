package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FillColor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9.525.626l.707.707l2.475 2.475l8.485 8.485l.708.707l-.707.707l-8.486 8.486L12 22.9l-.707-.707l-8.485-8.486L2.1 13l.707-.707l7.778-7.778l-1.768-1.768l-.707-.707L9.525.626ZM5.93 12h12.142l-6.07-6.07L5.928 12Zm12.142 2H5.93L12 20.071l6.071-6.07Zm2.679 3.39l.784.99l.53.67c.581.733.581 1.847 0 2.58a1.677 1.677 0 0 1-1.314.657c-.53 0-1-.26-1.314-.657c-.581-.733-.581-1.847 0-2.58l.53-.67l.784-.99Z"/>`),
		g.Group(children),
	)
}