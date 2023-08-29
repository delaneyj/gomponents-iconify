package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightningAltSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.82 3a.5.5 0 0 0-.847-.36l-1.778 1.71a35.551 35.551 0 0 0-6.63 8.715a.5.5 0 0 0 .435.746h4.31V21a.5.5 0 0 0 .837.37l.795-.725a35.498 35.498 0 0 0 7.001-8.78l.492-.87a.5.5 0 0 0-.435-.747h-4.18V3Z"/>`),
		g.Group(children),
	)
}