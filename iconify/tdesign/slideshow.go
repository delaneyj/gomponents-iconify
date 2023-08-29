package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Slideshow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 3h12v18H6V3Zm2 2v14h8V5H8ZM1 5h3.5v14H1v-2h1.5V7H1V5Zm18.5 0H23v2h-1.5v10H23v2h-3.5V5Z"/>`),
		g.Group(children),
	)
}