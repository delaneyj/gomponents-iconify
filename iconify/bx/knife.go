package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Knife(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.66 3.6a3 3 0 0 0-4.24 0l-.71.71l-7.07 7.07l2.12 2.12l-6.36 6.36l1.41 1.42L19.66 6.43c1.1-1.1 1.1-1.73.71-2.12z"/>`),
		g.Group(children),
	)
}