package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HourglassLowBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M204 75.64V40a20 20 0 0 0-20-20H72a20 20 0 0 0-20 20v36a20.1 20.1 0 0 0 8 16l48 36l-48 36a20.1 20.1 0 0 0-8 16v36a20 20 0 0 0 20 20h112a20 20 0 0 0 20-20v-35.64a20.13 20.13 0 0 0-7.94-16L147.9 128l48.16-36.4A20.13 20.13 0 0 0 204 75.64ZM166.31 172h-77L128 143ZM76 212v-16h104v16ZM180 73.65L128 113L76 74V44h104Z"/>`),
		g.Group(children),
	)
}