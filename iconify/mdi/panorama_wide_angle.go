package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanoramaWideAngle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 4c-4 0-6.8.6-9 1c-.5 2-1 3.9-1 7c0 3 .5 5 1 7c2.2.4 5 1 9 1s6.9-.6 9-1c.6-2 1-4 1-7s-.5-5.1-1-7c-2.1-.4-5-1-9-1Z"/>`),
		g.Group(children),
	)
}