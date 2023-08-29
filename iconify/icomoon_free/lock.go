package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M9.25 7H9V4c0-1.654-1.346-3-3-3H4C2.346 1 1 2.346 1 4v3H.75a.753.753 0 0 0-.75.75v7.5c0 .412.338.75.75.75h8.5c.412 0 .75-.338.75-.75v-7.5A.753.753 0 0 0 9.25 7zM3 4c0-.551.449-1 1-1h2c.551 0 1 .449 1 1v3H3V4z"/>`),
		g.Group(children),
	)
}