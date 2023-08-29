package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Icloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m26.682 14.302l.005-.036a8.463 8.463 0 0 0-15.901-4.036c-2.927-1.531-6.464.469-6.651 3.771A6.287 6.287 0 0 0 0 19.907a6.293 6.293 0 0 0 6.292 6.292h19.745c7.51-.021 8.115-11.057.646-11.896z"/>`),
		g.Group(children),
	)
}