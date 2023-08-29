package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UploadOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.11 21.46L2.39 1.73L1.11 3l5.45 5.45L5 10h3.11l.89.89V16h5.11l2 2H5v2h13.11l2.73 2.73l1.27-1.27M15 10h4l-7-7l-2.9 2.9l5.9 5.9V10Z"/>`),
		g.Group(children),
	)
}