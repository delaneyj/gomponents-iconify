package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GolfEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M4.05.638c-.518.543.223 1.175.678.675L5.224.77l2.332 1.393L3.306 4.5c-.293.162-.341.44-.234.721l.889 2.317l-.935 2.804c-.129.385.171.649.474.658a.468.468 0 0 0 .475-.341l.92-2.764l.853-.283l.252.505V10.5s0 .5.5.5s.5-.5.5-.5V8.117a.84.84 0 0 0-.064-.362l-1.281-3.34l2.552-1.403c.187-.08.29-.254.29-.512c0-.226-.217-.413-.456-.556L4.905.071c-.16-.096-.275-.043-.323.007l-.533.56zM3 2a1 1 0 1 0 0 2a1 1 0 0 0 0-2z" fill="currentColor"/>`),
		g.Group(children),
	)
}