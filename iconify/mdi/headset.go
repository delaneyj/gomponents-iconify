package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Headset(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 1c-5 0-9 4-9 9v7a3 3 0 0 0 3 3h3v-8H5v-2a7 7 0 0 1 7-7a7 7 0 0 1 7 7v2h-4v8h4v1h-7v2h6a3 3 0 0 0 3-3V10c0-5-4.03-9-9-9Z"/>`),
		g.Group(children),
	)
}