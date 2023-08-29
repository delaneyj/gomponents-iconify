package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LandmarkJpEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M6.75 3.5A1.25 1.25 0 1 1 5.5 2.25A1.25 1.25 0 0 1 6.75 3.5zm-4 2.75A1.25 1.25 0 1 0 4 7.5a1.25 1.25 0 0 0-1.25-1.25zm5.5 0A1.25 1.25 0 1 0 9.5 7.5a1.25 1.25 0 0 0-1.25-1.25z" fill="currentColor"/>`),
		g.Group(children),
	)
}