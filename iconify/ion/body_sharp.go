package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BodySharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<circle cx="256" cy="56" r="56" fill="currentColor"/><path fill="currentColor" d="M464 128H48v52h144l-32 325.13l51 6.87l21.65-192h47.02L301 512l51-6.98L320 180h144v-52z"/>`),
		g.Group(children),
	)
}