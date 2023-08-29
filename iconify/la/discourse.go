package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Discourse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16.11 3C8.993 3 3 8.716 3 15.773L3.006 29l13.103-.012C23.23 28.988 29 23.051 29 15.994C29 8.937 23.23 3 16.11 3zM16 8a8 8 0 0 1 0 16a7.96 7.96 0 0 1-3.908-1.023L8 24l1.121-3.93A7.946 7.946 0 0 1 8 16a8 8 0 0 1 8-8z"/>`),
		g.Group(children),
	)
}