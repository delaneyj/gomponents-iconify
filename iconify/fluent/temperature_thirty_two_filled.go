package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TemperatureThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 5.5a3.25 3.25 0 0 0-3.25 3.25v9.754l-.313.354a4.75 4.75 0 1 0 7.126 0l-.313-.354V8.75A3.25 3.25 0 0 0 16 5.5Zm-5.75 3.25a5.75 5.75 0 0 1 11.5 0v8.834a7.25 7.25 0 1 1-11.5 0V8.75Zm4.5 3.5a1.25 1.25 0 1 1 2.5 0v7.022a3 3 0 1 1-2.5 0V12.25Z"/>`),
		g.Group(children),
	)
}