package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CampsiteEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M5.92 2.19a.5.5 0 0 0-.85 0L1.5 8h-1a.42.42 0 0 0-.5.39V9.5a.49.49 0 0 0 .5.5h10a.49.49 0 0 0 .5-.5V8.39a.42.42 0 0 0-.5-.39h-1L5.92 2.19zM5.5 3l3 5h-6l3-5z" fill="currentColor"/>`),
		g.Group(children),
	)
}