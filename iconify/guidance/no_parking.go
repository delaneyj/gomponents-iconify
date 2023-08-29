package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoParking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M12.5 12.5h1a3.5 3.5 0 1 0 0-7H9V9m0 10v-7M3.868 3.868l16.264 16.264M12 23.5C5.649 23.5.5 18.351.5 12S5.649.5 12 .5S23.5 5.649 23.5 12S18.351 23.5 12 23.5Z"/>`),
		g.Group(children),
	)
}