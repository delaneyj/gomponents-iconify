package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PenOffSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m9 9.707l5.146 5.147a.5.5 0 0 0 .708-.708l-13-13a.5.5 0 1 0-.708.708L6.293 7L3.338 9.955a1.65 1.65 0 0 0-.398.644l-.914 2.743a.5.5 0 0 0 .632.632l2.743-.914a1.65 1.65 0 0 0 .644-.398L9 9.707Zm3.263-1.677l-1.056 1.056l.707.707l1.056-1.056a1.75 1.75 0 0 0 0-2.474L12.707 6l.733-.732a1.914 1.914 0 0 0-2.707-2.708L7.707 5.586l2.707 2.707L12 6.707l.263.263a.75.75 0 0 1 0 1.06Z"/>`),
		g.Group(children),
	)
}