package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Droplet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M13.51 7.393C12.483 4.527 10.305 1.953 8 0C5.695 1.953 3.518 4.527 2.49 7.393c-.635 1.772-.698 3.696.197 5.397C3.716 14.745 5.791 16 8 16s4.284-1.255 5.313-3.21c.895-1.701.832-3.624.197-5.397zm-1.967 4.466A4.045 4.045 0 0 1 8 14a4.03 4.03 0 0 1-2.377-.791c.207.027.416.041.627.041a5.056 5.056 0 0 0 4.428-2.676c.701-1.333.64-2.716.373-3.818c.227.44.42.878.576 1.311c.353.985.625 2.443-.084 3.791z"/>`),
		g.Group(children),
	)
}