package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LoaderCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="12" cy="20" r="2" fill="currentColor"/><circle cx="12" cy="4" r="2" fill="currentColor"/><circle cx="6.343" cy="17.657" r="2" fill="currentColor"/><circle cx="17.657" cy="6.343" r="2" fill="currentColor"/><circle cx="4" cy="12" r="2.001" fill="currentColor"/><circle cx="20" cy="12" r="2" fill="currentColor"/><circle cx="6.343" cy="6.344" r="2" fill="currentColor"/><circle cx="17.657" cy="17.658" r="2" fill="currentColor"/>`),
		g.Group(children),
	)
}