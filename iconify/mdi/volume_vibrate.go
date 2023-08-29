package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeVibrate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 9v6h4l5 5V4L8 9H4m12.55-6.53L15.5 3.53L17.93 6L15 9l2.93 3L15 15l2.93 3l-2.43 2.47l1.05 1.06L20 18l-2.93-3L20 12l-2.93-3L20 6l-3.45-3.53Z"/>`),
		g.Group(children),
	)
}