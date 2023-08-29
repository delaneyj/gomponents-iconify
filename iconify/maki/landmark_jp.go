package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LandmarkJp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M9.5 4a2 2 0 1 1-4 0a2 2 0 0 1 4 0zm-6 4a2 2 0 1 0 0 4a2 2 0 0 0 0-4zm8 0a2 2 0 1 0 0 4a2 2 0 0 0 0-4z"/>`),
		g.Group(children),
	)
}