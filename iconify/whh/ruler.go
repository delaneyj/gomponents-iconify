package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ruler(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m1015.43 855l-160 160q-10 9-23 9t-22-9l-63-63l81-81q9-9 0-17l-17-18q-9-9-18 0l-81 81l-125-125l81-81q9-9 0-17l-17-18q-9-9-18 0l-81 81l-125-125l81-81q9-9 0-17l-17-18q-9-9-18 0l-81 81l-125-125l81-81q9-9 0-17l-17-18q-9-9-18 0l-81 81l-125-125l81-81q9-9 0-17l-17-18q-9-9-18 0l-81 81l-63-63q-9-9-9-22t9-23l160-160q10-9 23-9t22 9l801 801q9 9 9 22t-9 23z"/>`),
		g.Group(children),
	)
}