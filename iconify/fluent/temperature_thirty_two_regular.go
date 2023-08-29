package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TemperatureThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 5a3 3 0 0 0-3 3v9.975l-.333.298a5 5 0 1 0 6.666 0L19 17.975V8a3 3 0 0 0-3-3Zm-5 3a5 5 0 0 1 10 0v9.101a7 7 0 1 1-10 0V8Zm4 4a1 1 0 1 1 2 0v7.17a3.001 3.001 0 1 1-2 0V12Z"/>`),
		g.Group(children),
	)
}