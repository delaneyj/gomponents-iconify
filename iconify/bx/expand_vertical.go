package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 19.24l-4.95-4.95l-1.41 1.42L12 22.07l6.36-6.36l-1.41-1.42L12 19.24zM5.64 8.29l1.41 1.42L12 4.76l4.95 4.95l1.41-1.42L12 1.93L5.64 8.29z"/>`),
		g.Group(children),
	)
}