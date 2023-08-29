package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pause(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 19V6h4v13h-4m-7 0V6h4v13H6M7 7v11h2V7H7m7 0v11h2V7h-2Z"/>`),
		g.Group(children),
	)
}