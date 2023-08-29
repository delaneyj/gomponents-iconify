package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArmchairFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M40 84.11V72a40 40 0 0 1 40-40h96a40 40 0 0 1 40 40v12.11a3.92 3.92 0 0 1-4 3.89a44.07 44.07 0 0 0-43.26 36H87.26A44.07 44.07 0 0 0 44 88a3.92 3.92 0 0 1-4-3.89ZM212 104a28 28 0 0 0-28 28v36a8 8 0 0 1-16 0v-28H88v28a8 8 0 0 1-16 0v-36a28 28 0 1 0-32 27.71V200a16 16 0 0 0 16 16h144a16 16 0 0 0 16-16v-40.29a28 28 0 0 0-4-55.71Z"/>`),
		g.Group(children),
	)
}