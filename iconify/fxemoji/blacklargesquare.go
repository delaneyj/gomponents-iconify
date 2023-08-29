package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blacklargesquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#132028" d="M510.396 31.453c0-16.417-13.432-29.849-29.849-29.849H28.134C13.359 1.604 1.27 13.693 1.27 28.468v455.398c0 14.775 12.089 26.864 26.864 26.864h452.413c16.417 0 29.849-13.432 29.849-29.849V31.453z"/>`),
		g.Group(children),
	)
}