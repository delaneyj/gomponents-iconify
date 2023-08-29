package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MeasurementTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.5 1.586L22.914 6L21.5 7.414l-2-2v13.172l2-2L22.914 18L18.5 22.414L14.086 18l1.414-1.414l2 2V5.414l-2 2L14.086 6L18.5 1.586ZM2 2h10v20H2V2Zm2 2v16h6V4H4Z"/>`),
		g.Group(children),
	)
}