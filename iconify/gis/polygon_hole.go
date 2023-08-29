package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PolygonHole(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M33.39 20.396a3.5 3.5 0 0 0-3.11 1.971L9.573 65.117c-1.026 2.121.299 4.629 2.631 4.979l75.002 11.261c2.575.385 4.653-2.073 3.844-4.544L76.787 33.374a3.498 3.498 0 0 0-2.432-2.291l-40.03-10.572a3.503 3.503 0 0 0-.934-.116Zm9.159 20.737l21.88 10.44l-1.072 10.025l-24.878-5.116Z" color="currentColor"/>`),
		g.Group(children),
	)
}