package bxl

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Firebase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.239 15.063L7.21 2.381a.453.453 0 0 1 .847-.145l2.12 3.979l-4.938 8.848zM19.24 18.14L17.363 6.469a.454.454 0 0 0-.766-.246L4.76 18.14l6.55 3.691c.411.23.912.23 1.323 0l6.607-3.691zM13.917 7.955L12.4 5.052a.452.452 0 0 0-.8 0L4.939 16.989l8.978-9.034z"/>`),
		g.Group(children),
	)
}