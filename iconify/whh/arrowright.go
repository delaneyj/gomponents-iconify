package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Arrowright(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m1005 558l-452 448q-19 19-45.5 19t-45.5-19l-92-91q-19-18-19-45t19-45l186-184H64q-27 0-45.5-19T0 577V449q0-27 18.5-45.5T64 385h493L370 200q-19-19-19-45.5t19-45.5l92-91q19-18 45.5-18T553 18l452 449q19 19 19 45.5t-19 45.5z"/>`),
		g.Group(children),
	)
}