package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Achievement(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M10.325 14.763A6.002 6.002 0 0 1 6 9V1h12v9M6 3H1v4c0 2.509 1.791 4 4 4h1m14.034-.115C21.742 10.49 23 9.103 23 7V3h-5m-8 16H5v4h11.5m0-13a6.5 6.5 0 1 0 0 13a6.5 6.5 0 0 0 0-13Zm3.5 4l-4.5 4.5L13 16m-2.794-.576A4 4 0 0 0 8 19"/>`),
		g.Group(children),
	)
}