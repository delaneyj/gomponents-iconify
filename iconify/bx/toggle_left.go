package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 9c-1.628 0-3 1.372-3 3s1.372 3 3 3s3-1.372 3-3s-1.372-3-3-3z"/><path fill="currentColor" d="M16 6H8c-3.3 0-5.989 2.689-6 6v.016A6.01 6.01 0 0 0 8 18h8a6.01 6.01 0 0 0 6-5.994V12c-.009-3.309-2.699-6-6-6zm0 10H8a4.006 4.006 0 0 1-4-3.99C4.004 9.799 5.798 8 8 8h8c2.202 0 3.996 1.799 4 4.006A4.007 4.007 0 0 1 16 16zm4-3.984l.443-.004l.557.004h-1z"/>`),
		g.Group(children),
	)
}