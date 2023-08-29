package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompassTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.625 3.133l-1.5 1.5A7.98 7.98 0 0 0 12 4c-4.42 0-8 3.58-8 8s3.58 8 8 8s8-3.58 8-8a7.979 7.979 0 0 0-.633-3.125l1.5-1.5A9.951 9.951 0 0 1 22 12c0 5.52-4.48 10-10 10S2 17.52 2 12S6.48 2 12 2c1.668 0 3.242.41 4.625 1.133Zm1.739 1.089l1.414 1.414L12 13.414L10.586 12l7.778-7.778Z"/>`),
		g.Group(children),
	)
}