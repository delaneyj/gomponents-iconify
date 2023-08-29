package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraRear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 6a2 2 0 0 1-2-2a2 2 0 0 1 2-2c1.09 0 2 .9 2 2a2 2 0 0 1-2 2m5-6H7a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V2a2 2 0 0 0-2-2m-3 20v2h5v-2m-9 0H5v2h5v2l3-3l-3-3v2Z"/>`),
		g.Group(children),
	)
}