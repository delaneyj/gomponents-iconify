package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThermometerOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11 7.8l-2-2V5c0-1.66 1.34-3 3-3s3 1.34 3 3v6.8L11.2 8H13V5c0-.55-.45-1-1-1s-1 .45-1 1v2.8m11.11 13.66L2.39 1.73L1.11 3L9 10.89V13a4.997 4.997 0 0 0-1 7a4.997 4.997 0 0 0 7 1c.84-.63 1.4-1.5 1.71-2.4l4.13 4.13l1.27-1.27Z"/>`),
		g.Group(children),
	)
}