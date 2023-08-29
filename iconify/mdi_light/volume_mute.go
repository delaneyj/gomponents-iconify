package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeMute(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 9h4l4-4h2v15h-2l-4-4H2V9m1 6h3.41l4 4H11V6h-.59l-4 4H3v5m10.96.33l2.83-2.83l-2.83-2.83l.71-.71l2.83 2.83l2.83-2.83l.71.71l-2.83 2.83l2.83 2.83l-.71.71l-2.83-2.83l-2.83 2.83l-.71-.71Z"/>`),
		g.Group(children),
	)
}