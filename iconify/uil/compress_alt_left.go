package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompressAltLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.5 5.5a1 1 0 0 0-1 1v1.59l-5.79-5.8a1 1 0 0 0-1.42 1.42l5.8 5.79H6.5a1 1 0 0 0 0 2h4a1 1 0 0 0 .38-.08a1 1 0 0 0 .54-.54a1 1 0 0 0 .08-.38v-4a1 1 0 0 0-1-1Zm11.21 14.79l-5.8-5.79h1.59a1 1 0 0 0 0-2h-4a1 1 0 0 0-.38.08a1 1 0 0 0-.54.54a1 1 0 0 0-.08.38v4a1 1 0 0 0 2 0v-1.59l5.79 5.8a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/>`),
		g.Group(children),
	)
}