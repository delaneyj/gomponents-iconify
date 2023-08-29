package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Play(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.4 12.5L9 18.38L8 19V6l10.4 6.5m-1.9 0L9 7.8v9.4l7.5-4.7Z"/>`),
		g.Group(children),
	)
}