package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ColorInvert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 1.586l6.01 6.01a8.502 8.502 0 0 1 1.905 9.114a7.525 7.525 0 0 1-1.917 2.92a8.502 8.502 0 0 1-2.63 1.784A8.503 8.503 0 0 1 5.99 7.595L12 1.586Zm3.157 17.704c.453-.258.927-.612 1.451-1.1a6.477 6.477 0 0 0 1.445-2.21a6.7 6.7 0 0 0 .381-2.387c-.032-1.742-.695-3.44-1.838-4.583L13 5.414v14.542a5.533 5.533 0 0 0 1.548-.368c.207-.088.41-.187.61-.298ZM11 5.414L7.404 9.01A6.5 6.5 0 0 0 11 20.03V5.414Z"/>`),
		g.Group(children),
	)
}