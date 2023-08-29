package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeSource(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 9v6h4l5 5V4L7 9H3m13 6h-2V9h2v6m4 4h-2V5h2v14Z"/>`),
		g.Group(children),
	)
}