package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Explore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22.707 9.293a1 1 0 0 0-1.023-.242l-9 3a1.001 1.001 0 0 0-.633.633l-3 9a1 1 0 0 0 1.265 1.265l9-3a1.001 1.001 0 0 0 .633-.633l3-9a1 1 0 0 0-.242-1.023ZM11.581 20.42l2.21-6.628l4.419 4.419Z"/><path fill="currentColor" d="M16 30a14 14 0 1 1 14-14a14.016 14.016 0 0 1-14 14Zm0-26a12 12 0 1 0 12 12A12.014 12.014 0 0 0 16 4Z"/>`),
		g.Group(children),
	)
}