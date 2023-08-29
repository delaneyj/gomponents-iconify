package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReligiousChristian(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M6 .955V4H3v3h3v8h3V7h3V4H9V1c0-1-.978-1-.978-1H6.99S6 0 6 .955z"/>`),
		g.Group(children),
	)
}