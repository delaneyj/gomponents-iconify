package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsDoughnutChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M13 6c2.507.423 4.577 2.493 5 5h4c-.471-4.717-4.283-8.529-9-9v4z" fill="currentColor"/><path d="M18 13c-.478 2.833-2.982 4.949-5.949 4.949c-3.309 0-6-2.691-6-6C6.051 8.982 8.167 6.478 11 6V2c-5.046.504-8.949 4.773-8.949 9.949c0 5.514 4.486 10 10 10c5.176 0 9.445-3.903 9.949-8.949h-4z" fill="currentColor"/>`),
		g.Group(children),
	)
}