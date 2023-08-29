package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayCircleStroke(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18ZM1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12Zm7.5-5.63L18.25 12L8.5 17.63V6.37Zm2 3.465v4.33L14.25 12L10.5 9.835Z"/>`),
		g.Group(children),
	)
}