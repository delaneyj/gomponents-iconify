package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Simplybuilt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M863 936q0-112-79.5-191.5T592 665t-191 79.5T322 936t79 191t191 79t191.5-79T863 936zm863-1q0-112-79-191t-191-79t-191.5 79t-79.5 191q0 113 79.5 192t191.5 79t191-79.5t79-191.5zm322-809v1348q0 44-31.5 75.5T1940 1581H108q-45 0-76.5-31.5T0 1474V126q0-44 31.5-75.5T108 19h431q44 0 76 31.5t32 75.5v161h754V126q0-44 32-75.5t76-31.5h431q45 0 76.5 31.5T2048 126z"/>`),
		g.Group(children),
	)
}