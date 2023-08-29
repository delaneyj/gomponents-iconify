package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToothFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M168 24H88a56 56 0 0 0-56 55.75c0 42.72 8 75.4 14.7 95.28c8.72 25.8 20.62 45.49 32.64 54a15.67 15.67 0 0 0 9.13 2.97a16.09 16.09 0 0 0 16-14.9c.85-11.52 5-49.11 23.53-49.11s22.68 37.59 23.53 49.11a16.09 16.09 0 0 0 9.18 13.36a15.69 15.69 0 0 0 15.95-1.41c12-8.53 23.92-28.22 32.64-54c6.7-19.9 14.7-52.58 14.7-95.3A56 56 0 0 0 168 24Zm3 56.57a8 8 0 1 1-6 14.85l-37-14.81l-37 14.81a8 8 0 1 1-6-14.85L106.46 72L85 63.42a8 8 0 1 1 6-14.85l37 14.81l37-14.81a8 8 0 1 1 6 14.85L149.54 72Z"/>`),
		g.Group(children),
	)
}