package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M1.5 0C.67 0 0 .67 0 1.5c0 .66.41 1.2 1 1.41V5.1c-.59.2-1 .75-1 1.41c0 .83.67 1.5 1.5 1.5S3 7.34 3 6.51c0-.6-.34-1.1-.84-1.34c.09-.09.21-.16.34-.16h2c.82 0 1.5-.68 1.5-1.5v-.59c.59-.2 1-.75 1-1.41C7 .68 6.33.01 5.5.01S4 .68 4 1.51c0 .66.41 1.2 1 1.41v.59c0 .28-.22.5-.5.5h-2c-.17 0-.35.04-.5.09V2.91c.59-.2 1-.75 1-1.41C3 .67 2.33 0 1.5 0z"/>`),
		g.Group(children),
	)
}