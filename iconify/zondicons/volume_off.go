package zondicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m15 8.59l-2.12-2.13l-1.42 1.42L13.6 10l-2.13 2.12l1.42 1.42L15 11.4l2.12 2.13l1.42-1.42L16.4 10l2.13-2.12l-1.42-1.42L15 8.6zM4 7H0v6h4l5 5V2L4 7z"/>`),
		g.Group(children),
	)
}