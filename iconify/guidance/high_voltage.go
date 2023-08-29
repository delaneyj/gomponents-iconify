package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HighVoltage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M3.5 14.25v.25h7.974A30.864 30.864 0 0 1 7.5 23v.5c1.32 0 3.184-.544 5.089-1.273c5.044-1.931 7.912-7.076 7.915-12.477L20.5 9.5h-5.586A25.22 25.22 0 0 0 16.5.699V.5h-9v.08a25.36 25.36 0 0 1-4 13.67Z"/>`),
		g.Group(children),
	)
}